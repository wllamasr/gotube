package app

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/wllamasr/golangtube/config/db"
)

var (
	router = gin.Default()
)

func init() {
	_ = godotenv.Load()
	db.LoadClient()
}

func StartApplication() {
	mapUrls()
	_ = router.Run(":8080")
}
