package account

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	CreateAccountService(createAccount CreateAccountRequest) (Account, error)
	LoginAccountService(loginAccount CreateAccountRequest) (Account, error)
	FindByIDAccountService(userID int) (Account, error)
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

func (s *service) LoginAccountService(loginAccount CreateAccountRequest) (Account, error) {
	findByEmail, err := s.repository.FindByEmailRepository(loginAccount.Email)
	if err != nil {
		return findByEmail, err
	} else {
		if findByEmail.ID == 0 {
			return findByEmail, errors.New("no acces account")
		}
		err := bcrypt.CompareHashAndPassword([]byte(findByEmail.Password), []byte(loginAccount.Password))
		if err != nil {
			return findByEmail, err
		}
		return findByEmail, nil
	}
}

func (s *service) FindByIDAccountService(userID int) (Account, error) {
	findByIDAccount, err := s.repository.FindByIDAccountRepository(userID)
	if err != nil {
		return findByIDAccount, err
	}

	if findByIDAccount.ID == 0 {
		return findByIDAccount, errors.New("not data id user")
	}

	return findByIDAccount, nil
}
