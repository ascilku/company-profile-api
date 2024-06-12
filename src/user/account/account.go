package account

import "time"

type Account struct {
	ID        int
	Email     string
	Password  string
	CreatedAT time.Time
	UpdatedAt time.Time
}
