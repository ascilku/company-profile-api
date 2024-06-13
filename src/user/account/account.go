package account

import (
	"company-profile-api/src/user/profile"
	"time"
)

type Account struct {
	ID        int
	Email     string
	Password  string
	Profile   []profile.Profile
	CreatedAt time.Time
	UpdatedAt time.Time
}
