package handlers

import (
	"gin-sample-app/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheckModel struct {
	AppName string	`json:"app-name"`
	Version string	`json:"version"`
}

func HealthCheck(ctx *gin.Context) {
	cfg := config.GetConfig()
	healthcheckModel := HealthCheckModel{
		AppName: cfg.AppInfo.Name,
		Version: cfg.AppInfo.Version,
	}
	ctx.JSON(http.StatusOK, healthcheckModel)
}