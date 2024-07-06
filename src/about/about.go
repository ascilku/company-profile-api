package about

import "time"

type About struct {
	ID          int
	AccountID   int
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
