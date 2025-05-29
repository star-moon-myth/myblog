package controllers

import (
    "github.com/gin-gonic/gin"
    "myblog/models"
    "net/http"
    "gorm.io/gorm"
)

func GetPosts(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var posts []models.Post
        if err := db.Find(&posts).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, posts)
    }
}

func Index(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var posts []models.Post
        db.Find(&posts)
        c.HTML(http.StatusOK, "index.html", gin.H{
            "Posts": posts,
        })
    }
}