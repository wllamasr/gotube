package users

import (
	"database/sql/driver"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

type role string

var validate *validator.Validate

const (
	user      role = "user"
	admin     role = "admin"
	moderator role = "moderator"
	support   role = "support"
)

func (p *role) Scan(value interface{}) error {
	*p = role(value.([]byte))
	return nil
}

func (p role) Value() (driver.Value, error) {
	return string(p), nil
}

type User struct {
	gorm.Model
	FirstName string `json:"firstname" validate:"required"`
	LastName  string `json:"lastname" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" gorm:"default:'asd'" validate:"required"`
	role      role   `json:"role" gorm:"default:'user'" sql:"type:ENUM('user', 'admin', 'moderator', 'support')"`
}

func (user *User) Validate() error {
	validate = validator.New()

	if error := validate.Struct(user); error != nil {
		return error
	}

	return nil
}
