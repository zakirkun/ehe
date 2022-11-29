package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouters(debugMode bool) http.Handler {
	if !debugMode {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Server Uptime"})
	})

	return router
}
