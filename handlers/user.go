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
		ctx.JSON(http.StatusBadRequest, "Bad request")
		return
	}

	email := ctx.Query("email")
	db.Where("email = ?", email).Find(&accounts)

	if len(accounts) != 1 {
		ctx.JSON(http.StatusNotFound, "Account not found")
		return
	}

	var account models.AccountEp

	marshalledVal, err := json.Marshal(accounts[0])
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Bad request")
		return
	}

	err = json.Unmarshal(marshalledVal, &account)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Bad request")
		return
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Bad request")
	}

	ctx.JSON(http.StatusOK, account)
}

func SetUser(ctx *gin.Context) {
	var account *models.Account
	if err := ctx.BindJSON(&account); err != nil {
		ctx.JSON(http.StatusBadRequest, "Bad request")
		return
	}

	db, err := dbconnectors.Connect()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Bad request")
		return
	}
	if err := db.Create(&account).Error; err !=nil {
		ctx.JSON(http.StatusBadRequest, "Account with email " + account.Email +" already exists")
		return
	}

	ctx.JSON(http.StatusOK, account)

	return
}
