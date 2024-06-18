package certification

import "gorm.io/gorm"

type Repository interface {
	CreateCertificationRep(certificate Certificate) (Certificate, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateCertificationRep(certificate Certificate) (Certificate, error) {
	err := r.db.Create(&certificate).Error
	if err != nil {
		return certificate, err
	}
	return certificate, nil
}
