package routers

import (
	"github.com/gin-gonic/gin"
	"redrock-ecommerce/handlers"
)

func SetupProductRoutes(r *gin.Engine) {
	productGroup := r.Group("/product")
	{
		productGroup.GET("/list", handlers.GetProductsList)
		productGroup.GET("/search", handlers.SearchProduct)
		productGroup.GET("/info/:product_id", handlers.GetProductInfo)
		productGroup.GET("/:type", handlers.GetProductsByType)
	}
}
