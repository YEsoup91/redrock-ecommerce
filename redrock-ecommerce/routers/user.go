package routers

import (
	"github.com/gin-gonic/gin"
	"redrock-ecommerce/handlers"
	"redrock-ecommerce/middleware"
)

func SetupUserRoutes(r *gin.Engine) {
	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", handlers.RegisterHandler)
		userGroup.POST("/login", handlers.LoginHandler)
		userGroup.GET("/info/:user_id", middleware.AuthMiddleware(), handlers.GetUserInfo)
		userGroup.PUT("/password", middleware.AuthMiddleware(), handlers.UpdateUserPassword)
		userGroup.PUT("/info", middleware.AuthMiddleware(), handlers.UpdateUserInfo)
	}
}
