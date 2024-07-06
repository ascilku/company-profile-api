package about

type formatter struct {
	Description string `json:"description"`
}

func Formatter(about About) formatter {
	formatter := formatter{
		Description: about.Description,
	}
	return formatter
}
