package main

import (
	"company-profile-api/database/connection_db"
	"company-profile-api/handler"
	"company-profile-api/src/user/account"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	connectionDB, err := connection_db.ConnectionDB()
	if err != nil {
		log.Fatal("failed connection to database ", err.Error())
	} else {
		// migration.MigrationAll(connectionDB)
		// connectionDB.Migrator().DropTable(&account.Account{})
		// connectionDB.Migrator().DropTable(&profile.Profile{})
		newRepository := account.NewRepository(connectionDB)
		newService := account.NewService(newRepository)
		newAccountHandler := handler.NewAccountHandler(newService)
		router := gin.Default()
		api := router.Group("api")
		api.POST("account", newAccountHandler.CreateAccountHandler)
		router.Run()
	}
}
