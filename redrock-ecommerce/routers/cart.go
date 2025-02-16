package routers

import (
	"github.com/gin-gonic/gin"
	"redrock-ecommerce/handlers"
	"redrock-ecommerce/middleware"
)

func SetupCartRoutes(r *gin.Engine) {
	cartGroup := r.Group("/cart")
	{
		cartGroup.POST("/add", middleware.AuthMiddleware(), handlers.AddToCart)
		cartGroup.GET("", middleware.AuthMiddleware(), handlers.GetCart)
	}
}
