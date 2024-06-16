package profile

type CreateProfileReq struct {
	AccountID int
	Name      string `binding:"required"`
}
