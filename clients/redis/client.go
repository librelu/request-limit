package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/request-limit/utils/utilerrors"
)

func NewClient(address, password string, db, maxRetries int, readTimeout, writeTimeout, dialTimeout time.Duration) (Handler, error) {
	if err := validateInput(address, db, maxRetries, readTimeout, writeTimeout, dialTimeout); err != nil {
		return nil, utilerrors.Wrap(err, "error when validateInput in NewClient")
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:         address,
		DB:           db,
		MaxRetries:   maxRetries,
		ReadTimeout:  readTimeout * time.Millisecond,
		WriteTimeout: writeTimeout * time.Millisecond,
		DialTimeout:  dialTimeout * time.Millisecond,
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, utilerrors.Wrap(err, "error when healthcheck redis in NewClient")
	}
	return &redisClient{
		client: rdb,
	}, nil
}

func validateInput(address string, db, maxRetries int, readTimeout, writeTimeout, dialTimeout time.Duration) error {
	if address == "" {
		return utilerrors.New(
			fmt.Sprintf("address can't be blank, current:%s", address),
		)
	}
	if db < 0 {
		return utilerrors.New(
			fmt.Sprintf("db can't be negative, current:%d", db),
		)
	}
	if maxRetries < 0 {
		return utilerrors.New(
			fmt.Sprintf("maxRetries can't be negative, current:%d", maxRetries),
		)
	}
	if readTimeout < 0 {
		return utilerrors.New(
			fmt.Sprintf("readTimeout can't be negative, current:%d", readTimeout),
		)
	}
	if writeTimeout < 0 {
		return utilerrors.New(
			fmt.Sprintf("writeTimeout can't be negative, current:%d", writeTimeout),
		)
	}
	if dialTimeout < 0 {
		return utilerrors.New(
			fmt.Sprintf("dialTimeout can't be negative, current:%d", dialTimeout),
		)
	}
	return nil
}
