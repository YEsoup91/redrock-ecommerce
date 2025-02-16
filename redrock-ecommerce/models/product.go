package models

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	gorm.Model
	Name        string    `gorm:"not null" json:"name"`
	Description string    `gorm:"not null" json:"description"`
	Type        string    `gorm:"not null" json:"type"`
	Price       float64   `gorm:"not null" json:"price"`
	Cover       string    `gorm:"not null" json:"cover"`
	Link        string    `gorm:"not null" json:"link"`
	PublishTime time.Time `gorm:"not null" json:"publish_time"`
	CommentNum  uint      `json:"comment_num"`
	IsAddedCart bool      `json:"is_added_cart"`
}
