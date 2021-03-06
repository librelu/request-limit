package trackers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/request-limit/clients/redis"
	"github.com/request-limit/utils/utilerrors"
)

// RateLimitMiddleware Block user when tries so many times in a period
func (t *Trackers) RateLimitMiddleware(c *gin.Context) {
	if err := validateRateLimitMiddlewareRequest(c); err != nil {
		c.AbortWithError(http.StatusBadRequest, utilerrors.Wrap(err, "[RateLimitMiddleware] can't pass the validation"))
		return
	}

	var (
		count int64
		err   error
	)
	key := getCacheKey(c.ClientIP())
	_, e := t.Redis.Get(c, key)
	switch e {
	case nil:
		if count, err = t.Redis.INCR(c, key); err != nil {
			c.AbortWithError(http.StatusInternalServerError, utilerrors.Wrap(err, "[RateLimitMiddleware] INCR from redis failed"))
			return
		}
	case redis.NotFoundError:
		if count, err = t.Redis.INCRAndExpire(c, key, t.ExpiredDuration); err != nil {
			c.AbortWithError(http.StatusInternalServerError, utilerrors.Wrap(err, "[RateLimitMiddleware] INCRAndExpire from redis failed"))
			return
		}
	default:
		c.AbortWithError(http.StatusInternalServerError, utilerrors.Wrap(e, "[RateLimitMiddleware] get data from redis failed"))
		return
	}
	if count > t.DefaultLimit {
		// Display Error plain text only as quiz expected.
		c.Data(http.StatusForbidden, "application/json; charset=utf-8", []byte("Error"))
		c.Abort()
		return
	}
}

func validateRateLimitMiddlewareRequest(c *gin.Context) error {
	if c.ClientIP() == "" {
		return utilerrors.New("can't provide a blank IP header")
	}
	return nil
}
