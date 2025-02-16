package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"redrock-ecommerce/config"
	"redrock-ecommerce/models"
)

// 获取商品列表
func GetProductsList(c *gin.Context) {
	db := config.GetDB()
	var products []models.Product
	db.Find(&products)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   products,
	})
}

// 搜索商品
func SearchProduct(c *gin.Context) {
	query := c.Query("query")
	db := config.GetDB()
	var products []models.Product

	// 模糊搜索
	db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", query)).Find(&products)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   products,
	})
}

// 获取商品详情
func GetProductInfo(c *gin.Context) {
	productID := c.Param("product_id")
	db := config.GetDB()
	var product models.Product

	if err := db.First(&product, productID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "商品不存在"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   product,
	})
}

// 获取分类下的商品列表
func GetProductsByType(c *gin.Context) {
	productType := c.Param("type")
	db := config.GetDB()
	var products []models.Product

	if err := db.Where("type = ?", productType).Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取商品列表"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   products,
	})
}
