package app

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/wllamasr/golangtube/config/database"
)

var (
	router = gin.Default()
)

func init() {
	_ = godotenv.Load()
	database.LoadClient()
}

func StartApplication() {
	mapUrls()
	router.Run(":8080")
}
