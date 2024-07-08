package experience

import (
	"company-profile-api/src/experience/skils"
	"company-profile-api/src/experience/tools"

	"gorm.io/gorm"
)

type Repository interface {
	CreateExperienceRep(experience Experience) (Experience, error)
	CreateSkilsRep(skill []skils.Skils) ([]skils.Skils, error)
	CreateToolsRep(tools []tools.Tools) ([]tools.Tools, error)
	FindByIDAccountRep(accountID int) (Experience, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateExperienceRep(experience Experience) (Experience, error) {
	err := r.db.Create(&experience).Error
	if err != nil {
		return experience, err
	}
	return experience, nil
}

func (r *repository) CreateSkilsRep(skill []skils.Skils) ([]skils.Skils, error) {
	err := r.db.Create(&skill).Error
	if err != nil {
		return skill, err
	}
	return skill, nil
}

func (r *repository) CreateToolsRep(tools []tools.Tools) ([]tools.Tools, error) {
	err := r.db.Create(&tools).Error
	if err != nil {
		return tools, err
	}
	return tools, nil
}

func (r *repository) FindByIDAccountRep(accountID int) (Experience, error) {
	var keyExperience Experience
	err := r.db.Where("account_id = ?", accountID).Find(&keyExperience).Error
	if err != nil {
		return keyExperience, err
	}
	return keyExperience, nil
}
