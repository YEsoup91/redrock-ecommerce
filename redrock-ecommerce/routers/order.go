package routers

import (
	"github.com/gin-gonic/gin"
	"redrock-ecommerce/handlers"
	"redrock-ecommerce/middleware"
)

func SetupOrderRoutes(r *gin.Engine) {
	orderGroup := r.Group("/order")
	{
		orderGroup.POST("", middleware.AuthMiddleware(), handlers.PlaceOrder)
	}
}
