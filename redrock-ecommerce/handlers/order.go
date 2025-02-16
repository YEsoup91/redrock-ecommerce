package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"redrock-ecommerce/config"
	"redrock-ecommerce/models"
)

// 下单
func PlaceOrder(c *gin.Context) {
	userID := c.MustGet("user").(*jwt.Token).Claims.(jwt.MapClaims)["user_id"].(float64)
	var order models.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order.UserID = uint(userID)
	order.Status = "pending" // 默认订单状态为 "pending"

	db := config.GetDB()
	if err := db.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "下单失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "success",
		"order_id": order.ID,
	})
}

// 获取用户订单列表
func GetUserOrders(c *gin.Context) {
	userID := c.MustGet("user").(*jwt.Token).Claims.(jwt.MapClaims)["user_id"].(float64)
	db := config.GetDB()
	var orders []models.Order

	if err := db.Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取订单"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   orders,
	})
}
