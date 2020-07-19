package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"os"
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
	Model
	FirstName                string    `json:"first_name" validate:"required" gorm:"not null"`
	LastName                 string    `json:"last_name" validate:"required" gorm:"not null"`
	Email                    string    `json:"email" validate:"required,email" gorm:"type:varchar(100);unique_index;not null"`
	Password                 string    `json:"password" gorm:"default:'asd'" validate:"required"`
	Role                     role      `json:"role" gorm:"default:'user';size:255;type:enum('user', 'admin','moderator', 'support')" validate:"oneof=user admin moderator support"`
	EmailConfirmed           bool      `json:"email_confirmed" gorm:"default:false"`
	EmailConfirmationToken   string    `json:"email_confirmation_token"`
	EmailConfirmationExpires time.Time `json:"email_confirmation_expires"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (user *User) VerifyPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func (user *User) AfterCreate() {
	user.sendVerificationMail()
}

func (user *User) BeforeSave() {
	user.hashPassword()
}

func (user *User) sendVerificationMail() {

}

func (user *User) hashPassword() {
	hashedPassword, _ := Hash(user.Password)
	user.Password = string(hashedPassword)

}

func (user *User) BeforeCreate() {
	user.EmailConfirmationToken = uuid.New().String()
	user.EmailConfirmationExpires = time.Now().AddDate(0, 0, 3)
}

func (user *User) CreateToken() (string, error) {

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user"] = user

	atClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))

	if err != nil {
		return "", err
	}

	return token, nil
}
