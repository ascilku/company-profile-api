package profile

import (
	"time"
)

type Profile struct {
	ID            int
	AccountID     int
	Name          string
	ProfileImages []ProfileImage
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type ProfileImage struct {
	ID        int
	ProfileID int
	ImgUrl    string
	CreatedAt time.Time
	Updated   time.Time
}
