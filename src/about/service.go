package about

type Service interface {
	CreateOrUpdateAboutServ(create CreateReq) (bool, string, error)
	FindAboutServ(accountID int) (About, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateOrUpdateAboutServ(create CreateReq) (bool, string, error) {
	findByIdAccount, err := s.repository.FindByIdAccountAboutRep(create.AccountID)
	if err != nil {
		return false, "", err
	}

	if findByIdAccount.ID == 0 {
		var keyAbout About
		keyAbout.Description = create.Description
		keyAbout.AccountID = create.AccountID
		_, err = s.repository.CreateAboutRep(keyAbout)
		if err != nil {
			return false, "", err
		}
		return true, "created", nil
	}

	findByIdAccount.Description = create.Description
	_, err = s.repository.UpdateAboutRep(findByIdAccount)
	if err != nil {
		return false, "", err
	}

	return true, "updated", nil
}

func (s *service) FindAboutServ(accountID int) (About, error) {
	findById, err := s.repository.FindByIdAccountAboutRep(accountID)
	if err != nil {
		return findById, err
	}
	return findById, nil
}
