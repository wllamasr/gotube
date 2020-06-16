package users

import (
	"database/sql/driver"
	"github.com/go-playground/validator/v10"
	"github.com/wllamasr/golangtube/models"
	"golang.org/x/crypto/bcrypt"
	"time"
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
	models.MainModel
	FirstName                string `json:"firstname" validate:"required"`
	LastName                 string `json:"lastname" validate:"required"`
	Email                    string `json:"email" validate:"required,email" gorm:"type:varchar(100);unique_index"`
	Password                 string `json:"password" gorm:"default:'asd'" validate:"required"`
	Role                     role   `json:"role" gorm:"default:'user';size:255"`
	EmailConfirmed           bool   `json:"email_confirmed" gorm:"default:false"`
	EmailConfirmationToken   string `json:"email_confirmation_token"`
	EmailConfirmationExpires time.Time
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (user *User) BeforeSave() error {
	hashedPassword, err := Hash(user.Password)

	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	return nil
}

func (user *User) BeforeCreate() error {
	user.EmailConfirmationExpires = time.Time.AddDate(0, 0, 3, 0)
	return nil
}
