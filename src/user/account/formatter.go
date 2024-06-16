package account

type formatter struct {
	Token string `json:"token"`
}

func Formatter(account Account, token string) formatter {
	formatter := formatter{
		Token: token,
	}
	return formatter
}
