package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	port := os.Getenv("API_PORT")
	if port == "" {
		port = "3000"
	}
	server.Run(":" + port)

}
