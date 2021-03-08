package uploads

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wllamasr/golangtube/config/db"
	"github.com/wllamasr/golangtube/models"
	"github.com/wllamasr/golangtube/utils/errors"
	"net/http"
	"os"
	"path/filepath"
)

func List(c *gin.Context) {
	var uploads []models.Upload

	if err := db.Client.Find(&uploads).Error; err != nil {
		restError := errors.InternalServerError(err.Error())
		c.JSON(restError.Status, restError)
		return
	}

	c.JSON(http.StatusOK, uploads)
}

func Get(c *gin.Context) {
	param := c.Param("upload_id")

	var upload models.Upload

	if err := db.Client.Where(param).First(&upload).Error; err != nil {
		restError := errors.InternalServerError(err.Error())
		c.JSON(restError.Status, restError)
		return
	}

	c.JSON(http.StatusOK, upload)
}

func Create(c *gin.Context) {

	file, err := c.FormFile("file")

	if err != nil {
		restErr := errors.BadRequestError("Error reading the file")
		c.JSON(restErr.Status, restErr)
		return
	}

	var upload models.Upload

	if err := c.ShouldBind(&upload); err != nil {
		restErr := errors.BadRequestError("Error in the form. Try sending again")
		c.JSON(restErr.Status, restErr)
		return
	}

	uuid_ := uuid.New().String()

	filename := filepath.Base(uuid_)

	upload.OriginalFileName = file.Filename

	upload.FileName = uuid_

	err = c.SaveUploadedFile(file, fmt.Sprintf("./uploads/videos/%s", filename))

	if err != nil {
		restErr := errors.InternalServerError(err.Error())
		c.JSON(restErr.Status, restErr)
		return
	}
	user_, _ := c.Get("user")

	user := user_.(models.User)

	upload.User = user

	upload.PublicURL = fmt.Sprintf("%s/uploads/videos/%s", os.Getenv("APP_URL"), uuid_)
	upload.OriginalFileName = file.Filename

	db.Client.Model(&user).Association("")

	if err := db.Client.Create(&upload).Error; err != nil {
		_ = os.Remove(fmt.Sprintf("./uploads/videos/%s", uuid_))
		restErr := errors.BadRequestError(err.Error())
		c.JSON(restErr.Status, restErr)
		return
	}

	fmt.Println(c.Request.TLS)

	c.JSON(http.StatusOK, upload)
}
