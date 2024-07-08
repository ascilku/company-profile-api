package experience

import (
	"company-profile-api/src/experience/skils"
	"company-profile-api/src/experience/tools"
	"time"
)

type Experience struct {
	ID             int
	AccountID      int
	Tools          []tools.Tools
	Skils          []skils.Skils
	Position       string
	EmploymentType string // Full-time, Part-time, Freelance, Contract, Internship
	CompanyName    string
	Location       string
	StartDate      time.Time
	EndDate        time.Time
	Description    string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
