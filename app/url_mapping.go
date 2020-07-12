package app

import (
	"github.com/wllamasr/golangtube/controllers/auth"
	"github.com/wllamasr/golangtube/controllers/ping"
	"github.com/wllamasr/golangtube/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.POST("/register", users.CreateUser)
	router.POST("/auth/login", auth.Login)
}
