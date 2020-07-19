package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wllamasr/golangtube/config/db"
	"github.com/wllamasr/golangtube/models"
	"github.com/wllamasr/golangtube/utils"
	"github.com/wllamasr/golangtube/utils/errors"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restError := errors.BadRequestError("Invalid body JSON")
		c.JSON(restError.Status, restError)
		return
	}

	if err := utils.Validator(user); err != nil {
		restError := errors.BadRequestError(err)
		c.JSON(restError.Status, restError)
		return
	}

	if err := db.Client.Create(&user).Error; err != nil {
		restError := errors.BadRequestError(err.Error())
		c.JSON(restError.Status, restError)
		return
	}

	c.JSON(http.StatusOK, user)
}

func ListUsers(c *gin.Context){
	var users []models.User

	if err := db.Client.Find(&users).Error; err != nil {
		fmt.Println(users)
		restErr := errors.BadRequestError(err.Error())
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, users)
}