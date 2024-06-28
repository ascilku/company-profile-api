package certification

import "time"

type Service interface {
	CreateCertificationServ(certificate CreateCertificateRequest) (bool, error)
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
