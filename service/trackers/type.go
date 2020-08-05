package trackers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/request-limit/clients/redis"
)

const (
	ipKeyPattern = "ip:%s"
)

type Trackers struct {
	Redis           redis.Handler
	ExpiredDuration time.Duration
	DefaultLimit    int64
}

type Handler interface {
	RateLimitMiddleware(c *gin.Context)
	Track(c *gin.Context)
}
