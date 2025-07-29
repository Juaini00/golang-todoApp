package utils

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func BuildResponse(status int, message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  status,
		"message": message,
		"data":    data,
	}
}

func BuildValidatorErrorResponse(status int, message string, err error) map[string]interface{} {
	var errorList []FieldError

	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		for _, fe := range validationErrors {
			errorList = append(errorList, FieldError{
				Field:   fe.Field(),
				Message: fe.Tag(),
			})
		}
	}

	return map[string]interface{}{
		"status": status,
		message:  message,
		"errors": errorList,
	}

}

func BuildErrorResponse(status int, message string) map[string]interface{} {
	return map[string]interface{}{
		"status":  status,
		"message": message,
	}
}
