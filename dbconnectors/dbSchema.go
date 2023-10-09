package dbconnectors

import (
	"gin-sample-app/models"
)

func DBSchemaMigrate() {
	db, err := Connect()
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&models.Movie{})
	db.AutoMigrate(&models.Account{})
}
