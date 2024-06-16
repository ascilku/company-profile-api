package profile

type Service interface {
	CreateOrUpdateProfileServ(createProfile CreateProfileReq) (Profile, string, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateOrUpdateProfileServ(createProfile CreateProfileReq) (Profile, string, error) {
	var keyProfile Profile
	findByAccountID, err := s.repository.findByAccountIDProfileRepo(createProfile.AccountID)
	if err != nil {
		return findByAccountID, "", err
	}
	if findByAccountID.ID != 0 {
		findByAccountID.Name = createProfile.Name
		updateProfile, err := s.repository.UpdateProfileRepo(findByAccountID)
		if err != nil {
			return updateProfile, "", err
		}
		return updateProfile, "update", nil
	}

	keyProfile.Name = createProfile.Name
	keyProfile.AccountID = createProfile.AccountID
	keyCreateProfile, err := s.repository.CreateProfileRepo(keyProfile)
	if err != nil {
		return keyCreateProfile, "", err
	}
	return keyCreateProfile, "create", nil
}
