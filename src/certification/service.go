package certification

import (
	"errors"
	"time"
)

type Service interface {
	CreateCertificationServ(certificate CreateCertificateRequest) (bool, error)
	FindAllCertificateServ(accountID int) ([]Certificate, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateCertificationServ(certificate CreateCertificateRequest) (bool, error) {
	var keyCertificate Certificate
	keyCertificate.AccountID = certificate.AccountID
	keyCertificate.NameCertificate = certificate.NameCertificate
	keyCertificate.Description = certificate.Description
	keyCertificate.FileCertificate = certificate.FileCertificate
	timeParse, err := time.Parse("2006-01-02", certificate.OutYear)
	if err != nil {
		return false, err
	}
	keyCertificate.OutYear = timeParse
	_, err = s.repository.CreateCertificationRep(keyCertificate)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *service) FindAllCertificateServ(accountID int) ([]Certificate, error) {
	findAllCertificate, err := s.repository.FindAllCertificate(accountID)
	if err != nil {
		return findAllCertificate, err
	}
	if len(findAllCertificate) == 0 {
		return findAllCertificate, errors.New("not data certificate")
	}
	return findAllCertificate, nil
}
