package certification

import "time"

type CreateCertificateRequest struct {
	NameCertificate string
	Description     string
	FileCertificate string
	OutYear         time.Time
}
