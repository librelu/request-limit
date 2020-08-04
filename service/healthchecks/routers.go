package healthchecks

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthChecksRegister(router *gin.RouterGroup) {
	router.GET("/healthcheck", HealthCheckRegistration)
}

func HealthCheckRegistration(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
