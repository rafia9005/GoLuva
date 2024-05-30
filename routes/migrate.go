package routes

import (
	"github.com/rafia9005/GoLuva/database"
)

func RunMigrate(dataModel interface{}) {
	database.DB.AutoMigrate(dataModel)
}
