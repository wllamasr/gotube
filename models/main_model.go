package models

import (
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

var validate *validator.Validate

type MainModel struct {
	gorm.Model
}

func (model *MainModel) Validate() error {
	validate = validator.New()

	if error := validate.Struct(model); error != nil {
		return error
	}

	return nil
}
