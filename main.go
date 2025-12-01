package main

import (
	"main/bootstrap"
	"main/config"
	"main/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	cfg := config.LoadConfig()
	gin.SetMode(cfg.GinMode)

	docs.SwaggerInfo.Title = "Todo API"
	docs.SwaggerInfo.Description = "Rest API in golang for a simple Todo app."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:" + cfg.Port
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}
	db, err := bootstrap.InitializeApp(cfg)

	if err != nil {
		panic(db)
	}

	server := gin.Default()

	if cfg.GinMode != "release" {
		server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	server.Run(":" + cfg.Port)

}
