package about

type Service interface {
	CreateAboutServ(create CreateReq) (bool, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateAboutServ(create CreateReq) (bool, error) {
	var keyAbout About
	keyAbout.AccountID = create.AccountID
	keyAbout.Description = create.Description
	_, err := s.repository.CreateAboutRep(keyAbout)
	if err != nil {
		return false, err
	}
	return true, nil
}
