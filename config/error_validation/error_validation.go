package error_validation

import "github.com/go-playground/validator/v10"

func ErrorValidation(err error) []string {
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}
	return errors
}
