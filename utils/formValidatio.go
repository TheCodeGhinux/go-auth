package utils

import (
	"github.com/go-playground/validator/v10"
)

func FormatValidationErrors(errs validator.ValidationErrors) map[string]string {
	errorMessages := make(map[string]string)
	for _, err := range errs {
		errorMessages[err.Field()] = err.Tag() // return field and its validation issue
	}
	return errorMessages
}
