package main

import (
    "github.com/gin-gonic/gin"
    "myblog/config"
    "myblog/routes"
)

func main() {
    db := config.InitDB()

    r := gin.Default()

    routes.SetupRoutes(r, db)

    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "My blog is alive! 🎉 DB connected!",
        })
    })

    r.Run(":8080")
}