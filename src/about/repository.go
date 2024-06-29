package about

import "gorm.io/gorm"

type Repository interface {
	CreateAboutRep(about About) (About, error)
	FindByIdAboutRep(accountID int) (About, error)
	UpdateAboutRep(about About) (About, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateAboutRep(about About) (About, error) {
	err := r.db.Create(&about).Error
	if err != nil {
		return about, err
	}
	return about, nil
}

func (r *repository) FindByIdAboutRep(accountID int) (About, error) {
	var keyAbout About
	err := r.db.Where("account_id = ?", accountID).Find(&keyAbout).Error
	if err != nil {
		return keyAbout, err
	}
	return keyAbout, nil
}

func (r *repository) UpdateAboutRep(about About) (About, error) {
	err := r.db.Save(&about).Error
	if err != nil {
		return about, err
	}
	return about, nil
}
