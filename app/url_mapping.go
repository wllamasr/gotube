package app

import (
	"github.com/wllamasr/golangtube/app/middleware"
	"github.com/wllamasr/golangtube/controllers/auth"
	"github.com/wllamasr/golangtube/controllers/ping"
	"github.com/wllamasr/golangtube/controllers/uploads"
	"github.com/wllamasr/golangtube/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.Static("/uploads", "./uploads")
	api := router.Group("/api")
	{
		authRoutes := api.Group("/auth")
		{

			authRoutes.POST("/login", auth.Login)
			authRoutes.POST("/register", users.CreateUser)
		}

		usersRoutes := api.Group("/users")
		{
			usersRoutes.GET("/", users.ListUsers)
		}
		uploadRoutes := api.Group("upload")
		{
			uploadRoutes.GET("/", uploads.List)
			uploadRoutes.GET("/:upload_id", uploads.Get)
		}

		api.Use(middleware.Authenticated())
		{
			uploadRoutes := api.Group("upload")
			{
				uploadRoutes.POST("/", uploads.Create)
			}
		}
	}
}
