package profile

import (
	"company-profile-api/src/user/account"
	"time"
)

type Profile struct {
	ID        int
	AccountID int
	Name      string
	Account   account.Account
	CreatedAT time.Time
	UpdatedAt time.Time
}
