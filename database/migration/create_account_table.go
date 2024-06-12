package migration

import (
	"company-profile-api/src/user/account"

	"gorm.io/gorm"
)

func CreateAccountTable(db *gorm.DB) error {
	err := db.Debug().AutoMigrate(&account.Account{})
	if err != nil {
		return err
	}
	return nil
}
