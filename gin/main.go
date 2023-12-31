package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	server.Run(":8081") // 监听并在 0.0.0.0:8080 上启动服务
}
