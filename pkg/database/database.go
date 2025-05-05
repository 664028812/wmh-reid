package database

import (
	"fmt"
	"log"
	"wmh/config"
	"wmh/internal/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

// InitDB 初始化数据库连接
func InitDB(cfg config.DatabaseConfig) error {
	var err error

	// 使用SQLite作为默认数据库
	DB, err = gorm.Open(sqlite.Open("wmh.db"), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("连接数据库失败: %v", err)
	}

	// 自动迁移数据库结构
	if err := autoMigrate(); err != nil {
		return fmt.Errorf("数据库迁移失败: %v", err)
	}

	log.Println("数据库初始化成功")
	return nil
}

// autoMigrate 自动迁移数据库结构
func autoMigrate() error {
	return DB.AutoMigrate(
		&model.User{},
		&model.Goal{},
		&model.Progress{},
		&model.Reminder{},
	)
}

// CloseDB 关闭数据库连接
func CloseDB() {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err == nil {
			sqlDB.Close()
		}
	}
}
