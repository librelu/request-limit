package trackers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/request-limit/clients/redis"
	"github.com/request-limit/utils/utilerrors"
)

func NewTrackers(redis redis.Handler, expiredDuration time.Duration, defaultLimit int64) (TrackersHandler, error) {
	if err := validateTrackersInput(redis, expiredDuration, defaultLimit); err != nil {
		return nil, utilerrors.Wrap(err, "can't pass the validation in NewTrackers")
	}
	return &Trackers{
		Redis:           redis,
		ExpiredDuration: expiredDuration * time.Second,
		DefaultLimit:    defaultLimit,
	}, nil
}

func validateTrackersInput(redis redis.Handler, expiredDuration time.Duration, defaultLimit int64) error {
	if redis == nil {
		return utilerrors.New("redis Handler can't be nil")
	}
	if expiredDuration <= 0 {
		return utilerrors.New("expiredDuration can't be zero or negative")
	}
	if defaultLimit <= 0 {
		return utilerrors.New("defaultLimit can't be zero or negative")
	}
	return nil
}

func (t *Trackers) Track(c *gin.Context) {
	if err := validateTrackRequest(c); err != nil {
		c.AbortWithError(http.StatusForbidden, utilerrors.Wrap(err, "request can't pass the validation in Track"))
		return
	}
	key := getCacheKey(c.ClientIP())
	data, err := t.Redis.Get(c, key)
	if err != nil {
		msg := fmt.Sprintf("data from redis error:%v", key)
		log.Println(utilerrors.Wrap(err, msg))
		c.JSON(http.StatusInternalServerError, msg)
	}
	c.JSON(http.StatusOK, gin.H{
		"tries": data,
	})
	return
}

func validateTrackRequest(c *gin.Context) error {
	if c.ClientIP() == "" {
		return utilerrors.New("can't provide a blank IP header")
	}
	return nil
}