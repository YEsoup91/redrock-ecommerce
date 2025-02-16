package routers

import (
	"github.com/gin-gonic/gin"
	"redrock-ecommerce/handlers"
	"redrock-ecommerce/middleware"
)

func SetupCommentRoutes(r *gin.Engine) {
	commentGroup := r.Group("/comment")
	{
		commentGroup.GET("/:product_id", handlers.GetComments)
		commentGroup.POST("/:product_id", middleware.AuthMiddleware(), handlers.PostComment)
		commentGroup.PUT("/:comment_id", middleware.AuthMiddleware(), handlers.UpdateComment)
		commentGroup.DELETE("/:comment_id", middleware.AuthMiddleware(), handlers.DeleteComment)
	}
}
