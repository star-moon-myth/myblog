package config

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
    "myblog/models"
)

func InitDB() *gorm.DB {
    dsn := "root:521334@tcp(127.0.0.1:3306)/blog_db?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // 自动迁移，创建/更新表结构
    db.AutoMigrate(&models.Post{})

    return db
}