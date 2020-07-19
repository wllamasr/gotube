package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"github.com/wllamasr/golangtube/models"
	"github.com/wllamasr/golangtube/utils/errors"
	"os"
	"strings"
)

func Authenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")

		if header == "" {
			responseUnauthorized(c)
			return
		}

		if strings.HasPrefix(header, "Bearer ") == false {
			responseUnauthorized(c)
			return
		}

		tokenString := strings.Split(header, " ")[1]

		token, err := jwt.Parse(tokenString, parseToken)

		if err != nil {
			fmt.Println("Hola")
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			var user models.User
			err := mapstructure.Decode(claims["user"], &user)

			if err != nil {
				responseUnauthorized(c)
				return
			}
			c.Set("user", user)
			c.Next()
		}
	}
}

func parseToken(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}

	secret := os.Getenv("ACCESS_SECRET")

	return []byte(secret), nil
}

func responseUnauthorized(c *gin.Context){
	restError := errors.UnauthorizedError()
	c.JSON(restError.Status, restError)
	c.Abort()
}