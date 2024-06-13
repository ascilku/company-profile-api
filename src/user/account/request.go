package account

type CreateAccountRequest struct {
	Email    string `binding:"required,email"`
	Password string `binding:"required"`
}
