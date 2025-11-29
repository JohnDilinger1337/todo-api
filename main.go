package main

import (
	"os"

	"main/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "3000"
	}

	mode := os.Getenv("GIN_MODE")
	if mode == "" {
		mode = "debug"
	}
	gin.SetMode(mode)

	docs.SwaggerInfo.Title = "Todo API"
	docs.SwaggerInfo.Description = "Rest API in golang for a simple Todo app."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:" + port
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	server := gin.Default()

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.Run(":" + port)

}
