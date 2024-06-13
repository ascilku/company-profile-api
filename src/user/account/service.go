package account

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	CreateAccountService(createAccount CreateAccountRequest) (Account, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateAccountService(createAccount CreateAccountRequest) (Account, error) {
	findByEmail, err := s.repository.FindByEmailRepository(createAccount.Email)
	if err != nil {
		return findByEmail, err
	} else {
		if findByEmail.ID != 0 {
			return findByEmail, errors.New("email ready exists")
		} else {

			var keyAccount Account
			pass, err := bcrypt.GenerateFromPassword([]byte(createAccount.Password), bcrypt.MinCost)
			if err != nil {
				return keyAccount, err
			} else {
				keyAccount.Email = createAccount.Email
				keyAccount.Password = string(pass)

				createAccount, err := s.repository.CreateAccountRepository(keyAccount)
				if err != nil {
					return createAccount, err
				}
				return createAccount, nil
			}
		}

	}

}
