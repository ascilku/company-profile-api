package respon

type respon struct {
	Meta meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type meta struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Value   interface{} `json:"value"`
}

func ResponJson(message string, code int, value interface{}, data interface{}) respon {
	meta := meta{
		Message: message,
		Code:    code,
		Value:   value,
	}

	respon := respon{
		Meta: meta,
		Data: data,
	}

	return respon
}
