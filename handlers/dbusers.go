package handlers

import (
	"net/http"

	"gin-sample-app/dbconnectors"

	"github.com/gin-gonic/gin"
)

func GetAllDbUsers(ctx *gin.Context) {
	var usersList []map[string]any = make([]map[string]any, 0)
	getAllDbUsers(&usersList)
	ctx.JSON(http.StatusOK, usersList)
}

func getAllDbUsers(usersList *[]map[string]any) error {
	db, err := dbconnectors.Connect()
	if err != nil {
		return err
	}

	db.Raw("select distinct user,authentication_string FROM mysql.user;").Scan(usersList)
	return nil
}

func GetDbUser(ctx *gin.Context) {
	var dbUser DbUser
	if err := ctx.BindUri(&dbUser); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}
	//checking if user already exists
	var existingUsersList []map[string]any = make([]map[string]any, 0)
	getAllDbUsers(&existingUsersList)
	for _, user := range existingUsersList {
		if user["user"] == dbUser.Username {
			ctx.JSON(http.StatusOK, user)
			return
		}
	}
	ctx.JSON(http.StatusBadRequest, "Username: "+dbUser.Username+" does not exists")
}
func CreateDbUser(ctx *gin.Context) {
	var dbUser DbUser
	if err := ctx.BindJSON(&dbUser); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}
	//checking if user already exists
	var existingUsersList []map[string]any = make([]map[string]any, 0)
	getAllDbUsers(&existingUsersList)
	for _, user := range existingUsersList {
		if user["user"] == dbUser.Username {
			ctx.JSON(http.StatusBadRequest, "Username: "+dbUser.Username+" already exists")
			return
		}
	}
	db, err := dbconnectors.Connect()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}
	sql := "CREATE USER '" + dbUser.Username + "'@'localhost' IDENTIFIED BY '" + dbUser.Password + "'"
	db.Exec(sql)
	ctx.JSON(http.StatusOK, "User: "+dbUser.Username+" successfully created")
}

func DeleteUser(ctx *gin.Context) {
	var dbUser DbUser
	if err := ctx.BindUri(&dbUser); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}
	//checking if user already exists
	var existingUsersList []map[string]any = make([]map[string]any, 0)
	getAllDbUsers(&existingUsersList)
	for _, user := range existingUsersList {
		if user["user"] == dbUser.Username {
			db, err := dbconnectors.Connect()
			if err != nil {
				ctx.JSON(http.StatusBadRequest, err.Error())
			}
			sql := "Drop user '" + dbUser.Username + "'@'localhost'"
			db.Exec(sql)
			ctx.JSON(http.StatusOK, "User: "+dbUser.Username+" successfully deleted")
			return
		}
	}
	ctx.JSON(http.StatusBadRequest, "Username: "+dbUser.Username+" does not exists")
}

type DbUser struct {
	Username string `json:"username" uri:"username"`
	Password string `json:"password"`
}
