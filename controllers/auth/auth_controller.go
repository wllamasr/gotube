package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/wllamasr/golangtube/config/db"
	"github.com/wllamasr/golangtube/models"
	"github.com/wllamasr/golangtube/utils"
	"github.com/wllamasr/golangtube/utils/dto"
	"github.com/wllamasr/golangtube/utils/errors"
	"net/http"
)

func Login(c *gin.Context) {
	var body dto.LoginBody
	if err := c.ShouldBindJSON(&body); err != nil {
		restError := errors.BadRequestError("Invalid body")
		c.JSON(restError.Status, restError)
		return
	}

	if err := utils.Validator(body); err != nil {
		restError := errors.BadRequestError(err)
		c.JSON(restError.Status, restError)
		return
	}
	var user models.User

	db.Client.Where("email = ?", body.Email).First(&user)

	if err := user.VerifyPassword(body.Password); err != nil {
		restError := errors.UnauthorizedError()
		c.JSON(restError.Status, restError)
		return
	}

	c.JSON(http.StatusNotImplemented, gin.H{
		"err": "no implemented",
	})
}
