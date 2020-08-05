package redis

import (
	"context"
	"time"

	"github.com/request-limit/utils/utilerrors"
)

func (r *redisClient) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *redisClient) INCRAndExpire(ctx context.Context, key string, expiration time.Duration) (int64, error) {
	incResult, err := r.client.Incr(ctx, key).Result()
	if err != nil {
		return 0, utilerrors.Wrap(err, "failed when incr")
	}

	expireResult := r.client.Expire(ctx, key, expiration)
	if err := expireResult.Err(); err != nil {
		r.client.Del(ctx, key)
		return 0, utilerrors.Wrap(err, "failed when expire")
	}
	return incResult, nil
}

func (r *redisClient) INCR(ctx context.Context, key string) (int64, error) {
	incResult, err := r.client.Incr(ctx, key).Result()
	if err != nil {
		return 0, utilerrors.Wrap(err, "failed when incr")
	}

	return incResult, nil
}
