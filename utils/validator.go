package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/wllamasr/golangtube/utils/custom_validations"
)

type ValidationError struct {
	Key   string `json:"key"`
	Error string `json:"error"`
}

//Validate struct
func Validator(model interface{}) interface{} {
	validate := validator.New()
	registerCustomValidations(validate)
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

func registerCustomValidations(validator *validator.Validate){
	_ = validator.RegisterValidation("unrepeated", custom_validations.ValidateUnrepeated)
}
