package connection_db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectionDB() (*gorm.DB, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/companyProfileDB?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	return db, nil
}
