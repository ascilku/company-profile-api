package certification

type CreateCertificateRequest struct {
	AccountID       int
	NameCertificate string `form:"nameCertificate" binding:"required"`
	Description     string `form:"description" binding:"required"`
	FileCertificate string
	OutYear         string `form:"outYear" binding:"required"`
}
