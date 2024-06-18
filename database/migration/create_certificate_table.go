package migration

import (
	"company-profile-api/src/certification"

	"gorm.io/gorm"
)

func CreateCertificate(db *gorm.DB) error {
	err := db.Debug().AutoMigrate(&certification.Certificate{})
	if err != nil {
		return err
	}
	return nil
}
