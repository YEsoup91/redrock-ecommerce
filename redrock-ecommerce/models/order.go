package models

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	UserID     uint      `gorm:"not null" json:"user_id"`     // 下单用户 ID
	ProductID  uint      `gorm:"not null" json:"product_id"`  // 商品 ID
	Quantity   int       `gorm:"not null" json:"quantity"`    // 购买数量
	TotalPrice float64   `gorm:"not null" json:"total_price"` // 总价
	Status     string    `gorm:"not null" json:"status"`      // 订单状态（如 "pending", "completed"）
	CreatedAt  time.Time `gorm:"not null" json:"created_at"`  // 创建时间
}
