package redis

import (
	"context"
	"time"

	"github.com/request-limit/utils/utilerrors"
)

func (r *redisClient) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *redisClient) INCRAndExpire(ctx context.Context, key string, expiration time.Duration) error {
	if err := r.client.Incr(ctx, key).Err(); err != nil {
		return utilerrors.Wrap(err, "failed when incr")
	}

	expireResult := r.client.Expire(ctx, key, expiration)
	if err := expireResult.Err(); err != nil {
		r.client.Del(ctx, key)
		return utilerrors.Wrap(err, "failed when expire")
	}
	return nil
}

func (r *redisClient) INCR(ctx context.Context, key string) error {
	if err := r.client.Incr(ctx, key).Err(); err != nil {
		return utilerrors.Wrap(err, "failed when incr")
	}
	return nil
}
