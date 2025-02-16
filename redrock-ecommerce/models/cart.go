package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID    uint `gorm:"not null" json:"user_id"`
	ProductID uint `gorm:"not null" json:"product_id"`
	Quantity  uint `gorm:"not null" json:"quantity"`
}
