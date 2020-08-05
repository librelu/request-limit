package trackers

import (
	"github.com/gin-gonic/gin"
)

func TrackersRegister(router *gin.RouterGroup, trackerHandler Handler) {
	router.GET("/track", trackerHandler.Track)
}
