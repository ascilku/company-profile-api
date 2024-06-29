package certification

import "gorm.io/gorm"

type Repository interface {
	CreateCertificationRep(certificate Certificate) (Certificate, error)
	FindAllCertificate(accountID int) ([]Certificate, error)
	FindByIDRep(ID int) (Certificate, error)
	DeleteOneCertificate(certificate Certificate) (Certificate, error)
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

func (r *repository) FindAllCertificate(accountID int) ([]Certificate, error) {
	var keyCertificate []Certificate
	err := r.db.Where("account_id = ?", accountID).Find(&keyCertificate).Error
	if err != nil {
		return keyCertificate, err
	}
	return keyCertificate, nil
}

func (r *repository) FindByIDRep(ID int) (Certificate, error) {
	var keyCertificate Certificate
	err := r.db.Where("id = ?", ID).Find(&keyCertificate).Error
	if err != nil {
		return keyCertificate, err
	}
	return keyCertificate, nil
}

func (r *repository) DeleteOneCertificate(certificate Certificate) (Certificate, error) {
	err := r.db.Delete(&certificate).Error
	if err != nil {
		return certificate, err
	}
	return certificate, nil
}
