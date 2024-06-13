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
	CreatedAt time.Time
	UpdatedAt time.Time
}
