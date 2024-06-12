package validation

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"strings"
)

func ValidationErrors(err error) map[string]string {

	var errorsList = map[string]string{}

	var validationErrors validator.ValidationErrors

	if errors.As(err, &validationErrors) {
		for _, fieldError := range validationErrors {
			errorsList[strings.ToLower(fieldError.Field())] = getErrorMsg(fieldError.Tag())
		}
	}
	return errorsList
}

func validationErrorMessages() map[string]string {
	return map[string]string{
		"required": "This field is required.",
		"email":    "The email format is wrong.",
	}
}

func getErrorMsg(tag string) string {
	return validationErrorMessages()[tag]
}
