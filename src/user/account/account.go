package account

import (
	"company-profile-api/src/about"
	"company-profile-api/src/certification"
	"company-profile-api/src/experience"
	"company-profile-api/src/user/profile"
	"time"
)

type Account struct {
	ID          int
	Email       string
	Password    string
	About       about.About
	Profile     []profile.Profile
	Certificate []certification.Certificate
	Experience  []experience.Experience
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
