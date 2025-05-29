package main

import (
    "github.com/gin-gonic/gin"
    "myblog/config"
    "myblog/routes"
)

func main() {
    db := config.InitDB()

    r := gin.Default()

    // 加载模板
    r.LoadHTMLGlob("templates/*")
    // 静态文件
    r.Static("/static", "./static")

    routes.SetupRoutes(r, db)

    r.Run(":8080")
}