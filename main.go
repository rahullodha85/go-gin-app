package main

import (
	"fmt"
	"gin-sample-app/routes"
	"net/http"
	"os"
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
	addr := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	s := &http.Server{
		Addr:           addr,
		Handler:        r,
	}
	err := s.ListenAndServe() // listen and serve on 0.0.0.0:8080
	if err != nil {
		fmt.Print("An error has occurred, details: %w", err)
	}
}