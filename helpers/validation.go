package helpers

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func Validate(payload interface{}) error {
	validate = validator.New()
	err := validate.Struct(payload)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}
		return err
	}

	return nil
}

func FormatValidationError(err error) []string {
	var errorMsg []string
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, conditon: %s", e.Field(), e.ActualTag())
			errorMsg = append(errorMsg, errorMessage)
		}
	} else {
		errorMsg = append(errorMsg, "Terjadi kesalahan")
	}
	return errorMsg
}
