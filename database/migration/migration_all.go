package migration

import "gorm.io/gorm"

func MigrationAll(db *gorm.DB) error {
	err := CreateAccountTable(db)
	if err != nil {
		return err
	}

	err = CreateProfileTable(db)
	if err != nil {
		return err
	}
	return nil
}
