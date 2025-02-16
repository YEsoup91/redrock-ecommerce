package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"redrock-ecommerce/config"
	"redrock-ecommerce/models"
	"redrock-ecommerce/routers"
)

func main() {
	// 初始化数据库连接
	dbConf, err := config.NewDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 自动迁移
	dbConf.DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Cart{}, &models.Order{}, &models.Comment{})

	// 初始化 Gin 路由
	router := gin.Default()

	//配置相关路由
	routers.SetupUserRoutes(router)
	routers.SetupProductRoutes(router)
	routers.SetupCartRoutes(router)
	routers.SetupOrderRoutes(router)
	routers.SetupCommentRoutes(router)

	// 启动服务器
	log.Println("Starting server at :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
