package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type role string

const (
	user      role = "user"
	admin     role = "admin"
	moderator role = "moderator"
	support   role = "support"
)

type User struct {
	gorm.Model
	FirstName                string `json:"firstname" validate:"required"`
	LastName                 string `json:"lastname" validate:"required"`
	Email                    string `json:"email" validate:"required,email" gorm:"type:varchar(100);unique_index"`
	Password                 string `json:"password" gorm:"default:'asd'" validate:"required"`
	Role                     role   `json:"role" gorm:"default:'user';size:255"`
	EmailConfirmed           bool   `json:"email_confirmed" gorm:"default:false"`
	EmailConfirmationToken   string `json:"email_confirmation_token"`
	EmailConfirmationExpires time.Time
	Uploads                  Upload
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (user *User) VerifyPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
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
	user.EmailConfirmationExpires = user.EmailConfirmationExpires.AddDate(0, 0, 3)
	return nil
}
