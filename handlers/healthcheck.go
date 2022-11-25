package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheckModel struct {
	AppName string
	Version string
}

func HealthCheck(ctx *gin.Context) {
	healthcheckModel := HealthCheckModel{
		AppName: "gin-sample-app",
		Version: "1.0.0",
	}
	ctx.JSON(http.StatusOK, healthcheckModel)
}