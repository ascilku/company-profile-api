package account

type formatter struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

func Formatter(account Account, token string) formatter {
	formatter := formatter{
		Email: account.Email,
		Token: token,
	}
	return formatter
}
