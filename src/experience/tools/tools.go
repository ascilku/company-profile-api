package tools

import "time"

type Tools struct {
	ID           int
	ExperienceID int `gorm:"constraint:OnDelete:CASCADE"`
	Title        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
