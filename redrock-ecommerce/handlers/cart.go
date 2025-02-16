package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"redrock-ecommerce/config"
	"redrock-ecommerce/models"
)

// 添加商品到购物车
func AddToCart(c *gin.Context) {
	var req struct {
		ProductID uint `json:"product_id"`
		Quantity  uint `json:"quantity"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := c.MustGet("user").(*models.User)

	db := config.GetDB()
	cart := models.Cart{
		UserID:    user.ID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}

	if err := db.Create(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "添加失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

// 获取用户购物车中的商品列表
func GetCart(c *gin.Context) {
	userID := c.MustGet("user").(*jwt.Token).Claims.(jwt.MapClaims)["user_id"].(float64)

	db := config.GetDB()
	var carts []models.Cart
	var products []models.Product

	//查询购物车中的商品信息
	result := db.Joins("INNER JOIN products ON carts.product_id = products.id").Where("carts.user_id = ?", uint(userID)).Find(&carts)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取购物车数据"})
		return
	}

	//根据商品ID查询商品详细信息
	for _, cart := range carts {
		var product models.Product
		db.First(&product, "id = ?", cart.ProductID)
		products = append(products, product)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"products": products,
		},
	})
}
