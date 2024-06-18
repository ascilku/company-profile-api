package profile

type CreateProfileReq struct {
	AccountID     int
	Name          string `form:"name" binding:"required"`
	ProfileImages string
}
