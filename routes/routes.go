package routes

import (
    "github.com/gin-gonic/gin"
    "myblog/controllers"
    "gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
    r.GET("/posts", controllers.GetPosts(db))
    r.GET("/", controllers.Index(db))
}