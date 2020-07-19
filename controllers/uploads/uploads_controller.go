package uploads

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Create(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"status": "Not implemented",
	})
}
