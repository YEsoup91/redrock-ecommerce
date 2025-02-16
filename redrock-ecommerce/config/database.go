package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

// 包含数据库连接实例
type Database struct {
	DB *gorm.DB
}

// 全局变量存储数据库连接实例
var db *gorm.DB

// 初始化数据库连接
func NewDatabase() (*Database, error) {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		return nil, fmt.Errorf("DB_DSN environment variable not set")
	}

	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	//将数据库连接实例存储到全局变量中
	db = gormDB

	return &Database{DB: gormDB}, nil
}

// 返回全局的数据库连接实例
func GetDB() *gorm.DB {
	if db == nil {
		log.Fatal("Database connection is not initialized. Call NewDatabase() first.")
	}
	return db
}
