package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"redrock-ecommerce/config"
	"redrock-ecommerce/models"
	"strconv"
)

// 获取商品评论
func GetComments(c *gin.Context) {
	productID := c.Param("product_id")
	db := config.GetDB()
	var comments []models.Comment

	if err := db.Where("product_id = ?", productID).Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取评论"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   comments,
	})
}

// 发布评论
func PostComment(c *gin.Context) {
	productID := c.Param("product_id")
	userID := c.MustGet("user").(*jwt.Token).Claims.(jwt.MapClaims)["user_id"].(float64)
	var comment models.Comment

	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment.UserID = uint(userID)
	comment.ProductID, _ = strconv.ParseUint(productID, 10, 64)

	db := config.GetDB()
	if err := db.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "评论失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   comment.ID,
	})
}

// 更新评论
func UpdateComment(c *gin.Context) {
	commentID := c.Param("comment_id")
	userID := c.MustGet("user").(*jwt.Token).Claims.(jwt.MapClaims)["user_id"].(float64)

	db := config.GetDB()
	var comment models.Comment
	if err := db.First(&comment, commentID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "评论不存在"})
		return
	}

	if uint(userID) != comment.UserID {
		c.JSON(http.StatusForbidden, gin.H{"error": "无权限修改"})
		return
	}

	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Save(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

// 删除评论
func DeleteComment(c *gin.Context) {
	commentID := c.Param("comment_id")
	userID := c.MustGet("user").(*jwt.Token).Claims.(jwt.MapClaims)["user_id"].(float64)

	db := config.GetDB()
	var comment models.Comment
	if err := db.First(&comment, commentID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "评论不存在"})
		return
	}

	if uint(userID) != comment.UserID {
		c.JSON(http.StatusForbidden, gin.H{"error": "无权限删除"})
		return
	}

	if err := db.Delete(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
