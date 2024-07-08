package experience

type CreateExperienceRequest struct {
	AccountID      int
	Tools          []string `binding:"required"`
	Skils          []string `binding:"required"`
	Position       string   `binding:"required"`
	EmploymentType string   `binding:"required"` // Full-time, Part-time, Freelance, Contract, Internship
	CompanyName    string   `binding:"required"`
	Location       string   `binding:"required"`
	StartDate      string   `binding:"required"`
	EndDate        string   `binding:"required"`
	Description    string   `binding:"required"`
}
