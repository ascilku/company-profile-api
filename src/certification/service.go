package certification

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
	keyCertificate.NameCertificate = certificate.NameCertificate
	keyCertificate.Description = certificate.Description
	keyCertificate.FileCertificate = certificate.FileCertificate
	keyCertificate.OutYear = certificate.OutYear
	_, err := s.repository.CreateCertificationRep(keyCertificate)
	if err != nil {
		return false, err
	}
	return true, nil
}
