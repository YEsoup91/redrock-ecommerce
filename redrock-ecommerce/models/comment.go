package models

import (
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	gorm.Model
	UserID      uint      `gorm:"not null" json:"user_id"`
	ProductID   uint64    `gorm:"not null" json:"product_id"`
	Content     string    `gorm:"not null" json:"content"`
	CreatedAt   time.Time `gorm:"not null" json:"publish_time"`
	PraiseCount int       `json:"praise_count"`
	IsPraised   int       `json:"is_praised"` // 0: 未处理，1: 点赞，2: 点踩
}
