package migration

import (
	"company-profile-api/src/user/profile"

	"gorm.io/gorm"
)

func CreateProfileTable(db *gorm.DB) error {
	err := db.Debug().AutoMigrate(&profile.Profile{})
	if err != nil {
		return err
	}
	return nil
}
