package skils

import "time"

type Skils struct {
	ID           int
	ExperienceID int `gorm:"constraint:OnDelete:CASCADE"`
	Title        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
