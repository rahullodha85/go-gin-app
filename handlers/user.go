package handlers

import (
	"encoding/json"
	"gin-sample-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gin-sample-app/dbconnectors"
)

func GetUser(ctx *gin.Context) {
	var accounts []models.Account
	db, err := dbconnectors.Connect()
	if err != nil {
		return
	}

	email := ctx.Query("email")
	db.Where("email = ?", email).Find(&accounts)

	if len(accounts) != 1 {
		ctx.JSON(http.StatusNotFound, "Account not found")
	}

	var account models.AccountEp

	marshalledVal, err := json.Marshal(accounts[0])
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Bad request")
	}

	err = json.Unmarshal(marshalledVal, &account)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Bad request")
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Bad request")
	}

	ctx.JSON(http.StatusOK, account)
}

func SetUser(ctx *gin.Context) {
	var account *models.Account
	if err := ctx.BindJSON(&account); err != nil {
		return
	}

	db, err := dbconnectors.Connect()
	if err != nil {
		return
	}
	db.Create(&account)

	ctx.JSON(http.StatusOK, account)

	return
}
