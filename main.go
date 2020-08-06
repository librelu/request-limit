package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/request-limit/clients/redis"
	"github.com/request-limit/service/healthchecks"
	"github.com/request-limit/service/trackers"
	"github.com/request-limit/utils/utilerrors"
)

type handlers struct {
	redisHandler   redis.Handler
	trackerHandler trackers.Handler
}

type clients struct {
}

func main() {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	handlers, err := initHandlers()
	if err != nil {
		panic(err)
	}
	initEndpoints(engine, handlers)
	port := "8080"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	engine.Run(fmt.Sprintf("0.0.0.0:%s", port))
}

func initEndpoints(engine *gin.Engine, h *handlers) {
	baseAPI := engine.Group("")
	groupAPI := engine.Group("/api")
	v1GroupAPI := groupAPI.Group("/v1")
	healthchecks.HealthChecksRegister(baseAPI)
	v1GroupAPI.Use(h.trackerHandler.RateLimitMiddleware)
	trackers.TrackersRegister(v1GroupAPI, h.trackerHandler)
}

func initHandlers() (*handlers, error) {
	redisClient, err := redis.NewClient(
		os.Getenv("REDIS_USERNAME"),
		os.Getenv("REDIS_ADDRESS"),
		os.Getenv("REDIS_PASSWORD"),
		0, 0, 3000, 3000, 3000,
	)
	if err != nil {
		return nil, utilerrors.Wrap(err, "can't new redis client in initHandlers")
	}
	trackerClient, err := trackers.NewTrackers(
		redisClient,
		60,
		60,
	)
	if err != nil {
		return nil, utilerrors.Wrap(err, "can't new redis client in initHandlers")
	}
	return &handlers{
		redisHandler:   redisClient,
		trackerHandler: trackerClient,
	}, nil
}
