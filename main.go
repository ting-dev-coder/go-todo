package main

import (
	"gin-todo/middleware"
	"gin-todo/repository"
	"gin-todo/routers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化資料庫
	repository.Init()
	
	r := gin.Default()

	r.Use(middleware.LoggerMiddleware())

	routers.InitRouters(r)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.Run(":8080") // http://localhost:8080

}
