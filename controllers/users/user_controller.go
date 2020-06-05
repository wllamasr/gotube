package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wllamasr/golangtube/config/database"
	"github.com/wllamasr/golangtube/models/users"
	"github.com/wllamasr/golangtube/utils/errors"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user users.User

	if error := c.ShouldBindJSON(&user); error != nil {
		restError := errors.BadRequestError("Invalid body JSON")
		c.JSON(restError.Status, restError)
		return
	}

	if error := user.Validate(); error != nil {
		fmt.Println(error.Error())
		restError := errors.BadRequestError(error.Error())
		c.JSON(restError.Status, restError)
		return
	}

	database.Client.Create(&user)

	c.JSON(http.StatusOK, user)
}
