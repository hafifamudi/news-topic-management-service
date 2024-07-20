package errors

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateStruct(input interface{}) error {
	return validate.Struct(input)
}

func FormatValidationError(err error) map[string]string {
	formattedErrors := make(map[string]string)
	if _, ok := err.(*validator.InvalidValidationError); ok {
		formattedErrors["error"] = err.Error()
		return formattedErrors
	}

	for _, err := range err.(validator.ValidationErrors) {
		formattedErrors[strings.ToLower(err.Field())] = err.Tag()
	}
	return formattedErrors
}
