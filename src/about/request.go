package about

type CreateReq struct {
	AccountID   int
	Description string `binding:"required"`
}
