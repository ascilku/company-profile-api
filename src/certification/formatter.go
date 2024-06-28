package certification

import "time"

type FormatterCertificate struct {
	NameCertificate string    `json:"nameCertificate"`
	Description     string    `json:"description"`
	FileCertificate string    `json:"fileCertificate"`
	OutYear         time.Time `json:"outYear"`
}

func FormatterCertificatee(certificate Certificate) FormatterCertificate {
	formatterCertificate := FormatterCertificate{
		NameCertificate: certificate.NameCertificate,
		Description:     certificate.Description,
		FileCertificate: certificate.FileCertificate,
		OutYear:         certificate.OutYear,
	}

	return formatterCertificate
}
