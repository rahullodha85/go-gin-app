package dbconnectors

import (
	"gin-sample-app/dbModels"
)

func DBSchemaMigrate() {
	db, err := Connect()
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&dbModels.Movie{})
}
