package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

//go:generate mockgen -source=./type.go -destination=mocks/mocks.go
type Handler interface {
	Get(ctx context.Context, key string) (string, error)
	INCRAndExpire(ctx context.Context, key string, expiration time.Duration) (int64, error)
	INCR(ctx context.Context, key string) (int64, error)
}

type redisClient struct {
	client *redis.Client
}

const NotFoundError = redis.Nil
