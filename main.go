package main

import (
	"main/bootstrap"
	"main/config"
	"main/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"main/api"
	"main/controller"
	"main/repository"
	"main/service"
)

var (
	userRepo *repository.UserRepository
	jwtSvc   *service.JWTService
	authSvc  *service.AuthService

	authCtrl *controller.AuthController
	authAPI  *api.AuthAPI
)

func main() {
	cfg := config.LoadConfig()
	db, err := bootstrap.InitializeApp(cfg)

	userRepo = repository.NewUserRepository(db)
	jwtSvc = service.NewJWTService(cfg.JWTSecret, cfg.AppName, cfg.JWTExpiresAt)
	authSvc = service.NewAuthService(userRepo, jwtSvc)
	authCtrl = controller.NewAuthController(authSvc)
	authAPI = api.NewAuthAPI(authCtrl)

	gin.SetMode(cfg.GinMode)

	docs.SwaggerInfo.Title = "Todo API"
	docs.SwaggerInfo.Description = "Rest API in golang for a simple Todo app."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:" + cfg.Port
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	if err != nil {
		panic(db)
	}

	server := gin.Default()

	if cfg.GinMode != "release" {
		server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	apiRoutes := server.Group(docs.SwaggerInfo.BasePath)
	{
		auth := apiRoutes.Group("/auth")
		{
			auth.POST("/register", authAPI.RegisterRoute)
			auth.POST("/login", authAPI.LoginRoute)
		}
	}

	server.Run(":" + cfg.Port)

}
