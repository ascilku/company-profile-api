package account

import (
	"company-profile-api/src/certification"
	"company-profile-api/src/user/profile"
	"time"
)

type Account struct {
	ID          int
	Email       string
	Password    string
	Profile     []profile.Profile
	Certificate []certification.Certificate
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
