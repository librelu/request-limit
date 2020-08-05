package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type Handler interface {
	Get(ctx context.Context, key string) (string, error)
	INCRAndExpire(ctx context.Context, key string, expiration time.Duration) error
	INCR(ctx context.Context, key string) error
}

type redisClient struct {
	client *redis.Client
}
