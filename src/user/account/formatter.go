package account

type formatter struct {
	Email string `json:"email"`
}

func Formatter(account Account) formatter {
	formatter := formatter{
		Email: account.Email,
	}
	return formatter
}
