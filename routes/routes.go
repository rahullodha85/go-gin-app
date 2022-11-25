package routes

import (
	"gin-sample-app/handlers"

	"github.com/gin-gonic/gin"
)

func AllRoutes(v1 *gin.RouterGroup) {
	router := v1.Group("v1")
	router.GET("/healthcheck", handlers.HealthCheck)
}