package certification

import "time"

type Certificate struct {
	ID              int
	AccountID       int
	NameCertificate string
	Description     string
	FileCertificate string
	OutYear         time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
