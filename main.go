package main

import (
	"github.com/gin-gonic/gin"
	"github.com/request-limit/service/healthchecks"
)

type clients struct {
}

func main() {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	// init clients
	// init endpoints
	initEndpoints(engine)
	// startup gin server
	// close connection if server down
	engine.Run("0.0.0.0:8000")
}

func initClients() *clients {
	return nil
}

func initEndpoints(engine *gin.Engine) {
	baseAPI := engine.Group("")
	groupAPI := engine.Group("/api")
	v1GroupAPI := groupAPI.Group("/v1")
	_ = v1GroupAPI
	healthchecks.HealthChecksRegister(baseAPI)
}
