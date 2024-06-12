package main

import (
	"company-profile-api/database/migration"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/companyProfileDB?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed connection to database ", err.Error())
	} else {
		migration.MigrationAll(db)
	}
}
