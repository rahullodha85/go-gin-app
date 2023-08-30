package main

import (
	"fmt"
	"gin-sample-app/config"
	"gin-sample-app/dbconnectors"
	"gin-sample-app/routes"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	if os.Getenv("DBMIGRATE") == "true" {
		dbconnectors.DBSchemaMigrate()
		fmt.Print("test")
		return
	}
	cfg := config.LoadConfig(os.Getenv("ENV"))
	fmt.Printf("%v", cfg.Server)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	group := r.Group("")
	routes.AllRoutes(group)
	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	s := &http.Server{
		Addr:    addr,
		Handler: r,
	}
	err := s.ListenAndServe() // listen and serve on 0.0.0.0:8080
	if err != nil {
		fmt.Print("An error has occurred, details: %w", err)
	}
}
