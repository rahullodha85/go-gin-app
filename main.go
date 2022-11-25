package main

import (
	"fmt"
	"gin-sample-app/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	group := r.Group("")
	routes.AllRoutes(group)
	err := r.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		fmt.Print("An error has occurred, details: %s", err)
	}
}