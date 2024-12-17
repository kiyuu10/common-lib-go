package types

import (
	"container/heap"
	"encoding/json"
)

type Heap[T any] struct {
	*HeapSlice[T]
}

func NewHeap[T any](data []T, less func(this, that T) bool) *Heap[T] {
	heapSlice := &HeapSlice[T]{
		data: data,
		less: less,
	}
	heap.Init(heapSlice)
	return &Heap[T]{HeapSlice: heapSlice}
}

func (h *Heap[T]) Push(x T) {
	heap.Push(h.HeapSlice, x)
}

func (h *Heap[T]) Pop() {
	heap.Pop(h.HeapSlice)
}

func (h *Heap[T]) Top() T {
	return h.data[0]
}

// HeapSlice is an implementation cointainer/heap
type HeapSlice[T any] struct {
	data []T
	less func(this, that T) bool // This is min heap, reverse less to greater if want to use max heap
}

func (h HeapSlice[T]) List() []T { return h.data }

func (h HeapSlice[T]) Len() int { return len(h.data) }

func (h HeapSlice[T]) Less(i, j int) bool {
	return h.less(h.data[i], h.data[j])
}

func (h HeapSlice[T]) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *HeapSlice[T]) Push(x interface{}) {
	h.data = append(h.data, x.(T))
}

func (h *HeapSlice[T]) Pop() interface{} {
	old := h.data
	n := len(old)
	x := old[n-1]
	h.data = old[0 : n-1]
	return x
}

func (h *HeapSlice[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(h.data)
}
