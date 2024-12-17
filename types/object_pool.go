package types

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"time"

	atomicobj "go.uber.org/atomic"
)

const (
	ObjPoolDefaultSize = 5
)

var (
	ObjPoolErrClosed      = errors.New("object pool: cannot be closed")
	ObjPoolErrPoolTimeout = errors.New("object pool: get timeout")

	vObjectPoolTimerPool = sync.Pool{
		New: func() any {
			t := time.NewTimer(time.Hour)
			t.Stop()
			return t
		},
	}
)

type ObjectPool[T any] interface {
	NewObject(context.Context) (*ObjectPoolItem[T], error)
	CloseObject(*ObjectPoolItem[T]) error

	Get(context.Context) (*ObjectPoolItem[T], error)
	Put(context.Context, *ObjectPoolItem[T])
	Remove(context.Context, *ObjectPoolItem[T], error)

	TotalCount() int
	IdleCount() int
	Stats() ObjPoolStats

	Close() error
}

type (
	// ObjPoolStats contains pool state information and accumulated stats.
	ObjPoolStats struct {
		HitCount     uint32
		MissCount    uint32
		TimeoutCount uint32

		TotalCount uint32
		IdleCount  uint32
		StaleCount uint32
	}
	ObjPoolOptions struct {
		PoolSize         int
		MinIdleCount     int
		MaxObjectAge     time.Duration
		PoolTimeout      time.Duration
		IdleTimeout      time.Duration
		IdleReapInterval time.Duration
	}
)

type ObjectPoolItem[T any] struct {
	useTime atomicobj.Int64
	value   T

	pooled     bool
	createTime time.Time
}

func newObjPoolObject[T any](value T) *ObjectPoolItem[T] {
	var (
		now = time.Now()
		obj = ObjectPoolItem[T]{
			value:      value,
			createTime: now,
		}
	)
	obj.setUseTime(now)
	return &obj
}

func (pi *ObjectPoolItem[T]) Value() T {
	return pi.value
}

func (pi *ObjectPoolItem[T]) getUseTime() time.Time {
	return time.Unix(pi.useTime.Load(), 0)
}

func (pi *ObjectPoolItem[T]) setUseTime(tm time.Time) {
	pi.useTime.Store(tm.Unix())
}

type SimpleObjectPool[T any] struct {
	creator   func(context.Context) (T, error)
	closer    func(obj T) error
	validator func(obj T) error
	opts      ObjPoolOptions

	createErrorCount atomicobj.Uint32
	lastCreateError  atomicobj.Error

	queueChn chan struct{}

	objectsMux  sync.Mutex
	objects     []*ObjectPoolItem[T]
	idleObjects []*ObjectPoolItem[T]
	size        int
	idleCount   int

	stats ObjPoolStats

	isClosed atomicobj.Bool
	closedCh chan struct{}
}

var _ ObjectPool[any] = (*SimpleObjectPool[any])(nil)

func NewSimpleObjectPool[T any](
	creator func(context.Context) (T, error),
	closer func(obj T) error,
	validator func(obj T) error,
	opts ObjPoolOptions,
) *SimpleObjectPool[T] {
	if opts.PoolSize == 0 {
		opts.PoolSize = ObjPoolDefaultSize
	}
	pool := &SimpleObjectPool[T]{
		creator:   creator,
		closer:    closer,
		validator: validator,
		opts:      opts,

		queueChn:    make(chan struct{}, opts.PoolSize),
		objects:     make([]*ObjectPoolItem[T], 0, opts.PoolSize),
		idleObjects: make([]*ObjectPoolItem[T], 0, opts.PoolSize),
		closedCh:    make(chan struct{}),
	}

	pool.objectsMux.Lock()
	pool.checkMinIdleObjects()
	pool.objectsMux.Unlock()

	if opts.IdleTimeout > 0 && opts.IdleReapInterval > 0 {
		go pool.reaper(opts.IdleReapInterval)
	}

	return pool
}

func (p *SimpleObjectPool[T]) checkMinIdleObjects() {
	if p.opts.MinIdleCount == 0 {
		return
	}
	for p.size < p.opts.PoolSize && p.idleCount < p.opts.MinIdleCount {
		p.size++
		p.idleCount++
		go func() {
			if err := p.addIdleObject(); err != nil {
				p.objectsMux.Lock()
				p.size--
				p.idleCount--
				p.objectsMux.Unlock()
			}
		}()
	}
}

func (p *SimpleObjectPool[T]) addIdleObject() error {
	obj, err := p.createObject(context.TODO(), true)
	if err != nil {
		return err
	}

	p.objectsMux.Lock()
	p.objects = append(p.objects, obj)
	p.idleObjects = append(p.idleObjects, obj)
	p.objectsMux.Unlock()
	return nil
}

func (p *SimpleObjectPool[T]) NewObject(ctx context.Context) (*ObjectPoolItem[T], error) {
	return p.newObject(ctx, false)
}

func (p *SimpleObjectPool[T]) newObject(ctx context.Context, pooled bool) (poolObj *ObjectPoolItem[T], err error) {
	obj, err := p.createObject(ctx, pooled)
	if err != nil {
		return
	}

	p.objectsMux.Lock()
	p.objects = append(p.objects, obj)
	if pooled {
		// If pool is full remove the object on next Put.
		if p.size >= p.opts.PoolSize {
			obj.pooled = false
		} else {
			p.size++
		}
	}
	p.objectsMux.Unlock()

	return obj, nil
}

func (p *SimpleObjectPool[T]) createObject(ctx context.Context, pooled bool) (_ *ObjectPoolItem[T], err error) {
	if p.isClosed.Load() {
		err = ObjPoolErrClosed
		return
	}

	if p.createErrorCount.Load() >= uint32(p.opts.PoolSize) {
		err = p.lastCreateError.Load()
		return
	}

	newValue, err := p.creator(ctx)
	if err != nil {
		p.lastCreateError.Store(err)
		p.createErrorCount.Inc()
		return
	}

	// internal.NewObjectectionsCounter.Add(ctx, 1)
	obj := newObjPoolObject(newValue)
	obj.pooled = pooled
	return obj, nil
}

// Get returns existed object from the pool or creates a new one.
func (p *SimpleObjectPool[T]) Get(ctx context.Context) (poolObj *ObjectPoolItem[T], err error) {
	if p.isClosed.Load() {
		err = ObjPoolErrClosed
		return
	}

	if err = p.waitTurn(ctx); err != nil {
		return
	}
	for {
		p.objectsMux.Lock()
		obj := p.popIdle()
		p.objectsMux.Unlock()
		if obj == nil {
			break
		}
		if p.isStaleObject(obj) {
			_ = p.CloseObject(obj)
			continue
		}
		atomic.AddUint32(&p.stats.HitCount, 1)
		return obj, nil
	}
	atomic.AddUint32(&p.stats.MissCount, 1)

	newObj, err := p.newObject(ctx, true)
	if err != nil {
		p.freeTurn()
		return nil, err
	}
	return newObj, nil
}

func (p *SimpleObjectPool[T]) getTurn() {
	p.queueChn <- struct{}{}
}

func (p *SimpleObjectPool[T]) waitTurn(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	select {
	case p.queueChn <- struct{}{}:
		return nil
	default:
		break
	}

	var timer = vObjectPoolTimerPool.Get().(*time.Timer)
	timer.Reset(p.opts.PoolTimeout)
	defer vObjectPoolTimerPool.Put(timer)

	select {
	case <-ctx.Done():
		if !timer.Stop() {
			<-timer.C
		}
		return ctx.Err()
	case p.queueChn <- struct{}{}:
		if !timer.Stop() {
			<-timer.C
		}
		return nil
	case <-timer.C:
		atomic.AddUint32(&p.stats.TimeoutCount, 1)
		return ObjPoolErrPoolTimeout
	}
}

func (p *SimpleObjectPool[T]) freeTurn() {
	<-p.queueChn
}

func (p *SimpleObjectPool[T]) popIdle() *ObjectPoolItem[T] {
	if len(p.idleObjects) == 0 {
		return nil
	}
	var (
		idx = len(p.idleObjects) - 1
		obj = p.idleObjects[idx]
	)
	p.idleObjects = p.idleObjects[:idx]
	p.idleCount--
	p.checkMinIdleObjects()
	return obj
}

func (p *SimpleObjectPool[T]) Put(ctx context.Context, obj *ObjectPoolItem[T]) {
	if p.validator != nil {
		err := p.validator(obj.value)
		if err != nil {
			p.Remove(ctx, obj, err)
			return
		}
	}

	if !obj.pooled {
		p.Remove(ctx, obj, nil)
		return
	}

	p.objectsMux.Lock()
	p.idleObjects = append(p.idleObjects, obj)
	p.idleCount++
	p.objectsMux.Unlock()
	p.freeTurn()
}

func (p *SimpleObjectPool[T]) Remove(ctx context.Context, obj *ObjectPoolItem[T], reason error) {
	p.removeObjectWithLock(obj)
	p.freeTurn()
	_ = p.closeObject(obj)
}

func (p *SimpleObjectPool[T]) CloseObject(obj *ObjectPoolItem[T]) error {
	p.removeObjectWithLock(obj)
	return p.closeObject(obj)
}

func (p *SimpleObjectPool[T]) removeObjectWithLock(obj *ObjectPoolItem[T]) {
	p.objectsMux.Lock()
	p.removeObject(obj)
	p.objectsMux.Unlock()
}

func (p *SimpleObjectPool[T]) removeObject(obj *ObjectPoolItem[T]) {
	for i, c := range p.objects {
		if c != obj {
			continue
		}
		p.objects = append(p.objects[:i], p.objects[i+1:]...)
		if obj.pooled {
			p.size--
			p.checkMinIdleObjects()
		}
		return
	}
}

func (p *SimpleObjectPool[T]) closeObject(obj *ObjectPoolItem[T]) error {
	if p.closer != nil {
		return p.closer(obj.value)
	}
	return nil
}

// TotalCount returns total number of objects.
func (p *SimpleObjectPool[T]) TotalCount() int {
	p.objectsMux.Lock()
	n := len(p.objects)
	p.objectsMux.Unlock()
	return n
}

// IdleCount returns number of idle objects.
func (p *SimpleObjectPool[T]) IdleCount() int {
	p.objectsMux.Lock()
	n := p.idleCount
	p.objectsMux.Unlock()
	return n
}

func (p *SimpleObjectPool[T]) Stats() ObjPoolStats {
	return ObjPoolStats{
		HitCount:     atomic.LoadUint32(&p.stats.HitCount),
		MissCount:    atomic.LoadUint32(&p.stats.MissCount),
		TimeoutCount: atomic.LoadUint32(&p.stats.TimeoutCount),

		TotalCount: uint32(p.TotalCount()),
		IdleCount:  uint32(p.IdleCount()),
		StaleCount: atomic.LoadUint32(&p.stats.StaleCount),
	}
}

// func (p *ObjectPool) Filter(fn func(T) bool) error {
// 	p.objectsMux.Lock()
// 	defer p.objectsMux.Unlock()

// 	var firstErr error
// 	for _, obj := range p.objects {
// 		if fn(obj) {
// 			if err := p.closeObject(obj); err != nil && firstErr == nil {
// 				firstErr = err
// 			}
// 		}
// 	}
// 	return firstErr
// }

func (p *SimpleObjectPool[T]) Close() error {
	if !p.isClosed.CompareAndSwap(false, true) {
		return ObjPoolErrClosed
	}
	close(p.closedCh)

	var firstErr error
	p.objectsMux.Lock()
	for _, obj := range p.objects {
		if err := p.closeObject(obj); err != nil && firstErr == nil {
			firstErr = err
		}
	}
	p.objects = nil
	p.size = 0
	p.idleObjects = nil
	p.idleCount = 0
	p.objectsMux.Unlock()

	return firstErr
}

func (p *SimpleObjectPool[T]) reaper(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// It is possible that ticker and closedCh arrive together,
			// and select pseudo-randomly pick ticker case, we double-check
			// here to prevent being executed after closed.
			if p.isClosed.Load() {
				return
			}
			if _, err := p.ReapStaleObjects(); err != nil {
				continue
			}
		case <-p.closedCh:
			return
		}
	}
}

func (p *SimpleObjectPool[T]) ReapStaleObjects() (int, error) {
	var n int
	for {
		p.getTurn()

		p.objectsMux.Lock()
		obj := p.reapStaleObject()
		p.objectsMux.Unlock()
		p.freeTurn()

		if obj != nil {
			_ = p.closeObject(obj)
			n++
		} else {
			break
		}
	}
	atomic.AddUint32(&p.stats.StaleCount, uint32(n))
	return n, nil
}

func (p *SimpleObjectPool[T]) reapStaleObject() *ObjectPoolItem[T] {
	if len(p.idleObjects) == 0 {
		return nil
	}

	obj := p.idleObjects[0]
	if !p.isStaleObject(obj) {
		return nil
	}

	p.idleObjects = append(p.idleObjects[:0], p.idleObjects[1:]...)
	p.idleCount--
	p.removeObject(obj)

	return obj
}

func (p *SimpleObjectPool[T]) isStaleObject(obj *ObjectPoolItem[T]) bool {
	if p.opts.IdleTimeout == 0 && p.opts.MaxObjectAge == 0 {
		return false
	}

	now := time.Now()
	if p.opts.IdleTimeout > 0 && now.Sub(obj.getUseTime()) >= p.opts.IdleTimeout {
		return true
	}
	if p.opts.MaxObjectAge > 0 && now.Sub(obj.createTime) >= p.opts.MaxObjectAge {
		return true
	}

	return false
}
