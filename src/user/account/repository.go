package account

import "gorm.io/gorm"

type Repository interface {
	CreateAccountRepository(account Account) (Account, error)
	FindByEmailRepository(email string) (Account, error)
	FindByIDAccountRepository(userID int) (Account, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateAccountRepository(account Account) (Account, error) {
	err := r.db.Create(&account).Error
	if err != nil {
		return account, err
	}
	return account, nil
}

func (r *repository) FindByEmailRepository(email string) (Account, error) {
	var keyAccount Account
	err := r.db.Where("email = ?", email).Find(&keyAccount).Error
	if err != nil {
		return keyAccount, err
	}
	return keyAccount, nil
}

func (r *repository) FindByIDAccountRepository(userID int) (Account, error) {
	var keyAccount Account
	err := r.db.Where("id = ?", userID).Find(&keyAccount).Error
	if err != nil {
		return keyAccount, err
	}
	return keyAccount, nil
}
