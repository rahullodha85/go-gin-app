package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context) {
	abc := Get("abc")
	ctx.JSON(http.StatusOK, abc)
}

func SetUser(ctx *gin.Context) {
	var abc map[string]interface{}
	if err := ctx.BindJSON(&abc); err != nil {
		return
	}
	Set("abc", abc)
	ctx.JSON(http.StatusOK, abc)
}


