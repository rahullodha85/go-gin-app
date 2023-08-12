package routes

import (
	"gin-sample-app/handlers"

	"github.com/gin-gonic/gin"
)

func AllRoutes(v1 *gin.RouterGroup) {
	router := v1.Group("v1")
	router.GET("/healthcheck", handlers.HealthCheck)
	router.POST("/user", handlers.SetUser)
	router.GET("/user", handlers.GetUser)
	router.GET("/dbusers", handlers.GetAllDbUsers)
	router.GET("/dbuser/:username", handlers.GetDbUser)
	router.POST("/dbuser", handlers.CreateDbUser)
	router.DELETE("/dbuser/:username", handlers.DeleteUser)
}
