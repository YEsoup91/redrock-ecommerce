package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username     string    `gorm:"uniqueIndex;not null" json:"username"`
	Password     string    `gorm:"not null" json:"password"`
	Nickname     string    `gorm:"not null" json:"nickname"`
	Email        string    `gorm:"uniqueIndex" json:"email"`
	Phone        string    `gorm:"uniqueIndex" json:"phone"`
	QQ           string    `gorm:"uniqueIndex" json:"qq"`
	Gender       string    `gorm:"not null" json:"gender"`
	Introduction string    `gorm:"not null" json:"introduction"`
	Birthday     time.Time `gorm:"not null" json:"birthday"`
}
