package profile

import "gorm.io/gorm"

type Repository interface {
	CreateProfileRepo(profile Profile) (Profile, error)
	UpdateProfileRepo(profile Profile) (Profile, error)
	findByIDAccountProfileRepo(accountID int) (Profile, error)
	findByAccountIDProfileRepo(accountID int) (Profile, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateProfileRepo(profile Profile) (Profile, error) {
	err := r.db.Create(&profile).Error
	if err != nil {
		return profile, err
	}
	return profile, nil
}

func (r *repository) findByIDAccountProfileRepo(accountID int) (Profile, error) {
	var keyProfile Profile
	err := r.db.Where("account_id = ?", accountID).Find(&keyProfile).Error
	if err != nil {
		return keyProfile, err
	}
	return keyProfile, nil
}

func (r *repository) UpdateProfileRepo(profile Profile) (Profile, error) {
	err := r.db.Save(&profile).Error
	if err != nil {
		return profile, err
	}
	return profile, nil
}

func (r *repository) findByAccountIDProfileRepo(accountID int) (Profile, error) {
	var keyProfile Profile
	err := r.db.Where("account_id = ?", accountID).Find(&keyProfile).Error
	if err != nil {
		return keyProfile, err
	}
	return keyProfile, nil
}
