package utils

import (
	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Key   string `json:"key"`
	Error string `json:"error"`
}

//Validate struct
func Validator(model interface{}) interface{} {
	validate := validator.New()

	var errors []ValidationError

	if err := validate.Struct(model); err != nil {
		//Create a new Array of ValidationErrors
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, ValidationError{
				Key:   err.Field(),
				Error: err.Tag(),
			})
		}
		return errors
	}
	return nil
}
