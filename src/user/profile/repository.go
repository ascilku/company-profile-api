package profile

import "gorm.io/gorm"

type Repository interface {
	// this is profile
	CreateProfileRepo(profile Profile) (Profile, error)
	UpdateProfileRepo(profile Profile) (Profile, error)
	FindByIDAccountProfileRepo(accountID int) (Profile, error)
	FindByAccountIDProfileRepo(accountID int) (Profile, error)
	// this is profile image
	CreateProfileImageRepo(profileImage ProfileImage) (ProfileImage, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// this is profile
func (r *repository) CreateProfileRepo(profile Profile) (Profile, error) {
	err := r.db.Create(&profile).Error
	if err != nil {
		return profile, err
	}
	return profile, nil
}

func (r *repository) FindByIDAccountProfileRepo(accountID int) (Profile, error) {
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

func (r *repository) FindByAccountIDProfileRepo(accountID int) (Profile, error) {
	var keyProfile Profile
	err := r.db.Where("account_id = ?", accountID).Find(&keyProfile).Error
	if err != nil {
		return keyProfile, err
	}
	return keyProfile, nil
}

// this is profile image
func (r *repository) CreateProfileImageRepo(profileImage ProfileImage) (ProfileImage, error) {
	err := r.db.Create(&profileImage).Error
	if err != nil {
		return profileImage, err
	}
	return profileImage, nil
}
