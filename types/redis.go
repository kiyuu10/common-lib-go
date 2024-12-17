package types

import (
	"context"
	"encoding/json"
	"errors"
	"net/url"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/go-redis/redis/v8"

	"gitea.alchemymagic.app/snap/go-common/erroy"
)

type RedisOptions struct {
	Prefix   string `json:"prefix"`
	DB       int    `json:"db,string"`
	Username string `json:"username"`
	Password string `json:"password"`

	DialTimeout  TimeDuration `json:"dial_timeout"`
	ReadTimeout  TimeDuration `json:"read_timeout"`
	WriteTimeout TimeDuration `json:"write_timeout"`

	PoolSize     int          `json:"pool_size,string"`
	MinIdleConns int          `json:"min_idle_conns,string"`
	MaxConnAge   TimeDuration `json:"max_conn_age"`
	PoolTimeout  TimeDuration `json:"pool_timeout"`
	IdleTimeout  TimeDuration `json:"idle_timeout"`
}

func NewRedisClient(opts redis.Options) *redis.Client {
	var client = redis.NewClient(&opts)
	client.AddHook(RedisSentryHook{})
	return client
}

func NewRedisClientWithDSN(dsnURL *url.URL) (_ *redis.Client, _ RedisOptions, err error) {
	ourOpts, err := ParseRedisOptions(dsnURL.Query())
	if err != nil {
		return
	}
	var (
		redisOpts = redis.Options{
			Addr:     dsnURL.Host,
			Username: ourOpts.Username,
			Password: ourOpts.Password,
			DB:       ourOpts.DB,

			DialTimeout:  ourOpts.DialTimeout.Duration,
			ReadTimeout:  ourOpts.ReadTimeout.Duration,
			WriteTimeout: ourOpts.WriteTimeout.Duration,

			PoolSize:     ourOpts.PoolSize,
			MinIdleConns: ourOpts.MinIdleConns,
			MaxConnAge:   ourOpts.MaxConnAge.Duration,
			PoolTimeout:  ourOpts.PoolTimeout.Duration,
			IdleTimeout:  ourOpts.IdleTimeout.Duration,
		}
		client = NewRedisClient(redisOpts)
	)
	return client, ourOpts, nil
}

func ParseRedisOptions(values url.Values) (_ RedisOptions, err error) {
	var (
		queryMap = urlQueryToMap(values)
		options  RedisOptions
	)
	queryJSON, err := json.Marshal(queryMap)
	if err != nil {
		err = erroy.WrapStack(err, "redis: encode dsn query")
		return
	}
	if err = json.Unmarshal(queryJSON, &options); err != nil {
		err = erroy.WrapStack(err, "redis: decode dsn query")
		return
	}
	return options, nil
}

type RedisSentryHook struct{}

func (rsh RedisSentryHook) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {
	cmdName := cmd.Name()
	if cmdName == "" {
		return ctx, nil
	}
	var (
		args = cmd.Args()
		span = sentry.StartSpan(ctx, "redis."+cmdName)
	)
	if len(args) > 1 {
		key := args[1]
		span.Data = map[string]any{
			"key": key,
		}
		if len(args) > 2 {
			for _, arg := range args[2:] {
				if timeout, ok := arg.(time.Duration); ok {
					span.Data["timeout"] = timeout.String()
					break
				}
			}
		}
	}
	return context.WithValue(span.Context(), cContextKeySentrySpan, span), nil
}

func (rsh RedisSentryHook) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	span, ok := ctx.Value(cContextKeySentrySpan).(*sentry.Span)
	if !ok {
		return nil
	}
	if err := cmd.Err(); err != nil {
		if errors.Is(err, redis.Nil) {
			span.Status = sentry.SpanStatusNotFound
		} else {
			span.Status = sentry.SpanStatusInternalError
		}
	} else {
		span.Status = sentry.SpanStatusOK
	}
	span.Finish()
	return nil
}

func (rsh RedisSentryHook) BeforeProcessPipeline(ctx context.Context, _ []redis.Cmder) (
	context.Context, error,
) {
	return ctx, nil
}

func (rsh RedisSentryHook) AfterProcessPipeline(_ context.Context, _ []redis.Cmder) error {
	return nil
}
