package controller

import (
	"main/config"
	domainErr "main/domain/error"
	"main/dto"
	"main/service"

	"time"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService *service.AuthService
	Cfg         *config.Config
}

func NewAuthController(authService *service.AuthService, cfg *config.Config) *AuthController {
	return &AuthController{AuthService: authService, Cfg: cfg}
}

func (c *AuthController) Register(ctx *gin.Context) error {
	var input dto.RegisterInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return err
	}

	user, err := c.AuthService.Register(input.Username, input.Email, input.Password)

	if err != nil {
		return err
	}

	token, err := c.AuthService.Login(user.Username, user.Password)
	if err != nil {
		return err
	}

	duration, err := time.ParseDuration(c.Cfg.JWTExpiresAt)
	if err != nil {
		return domainErr.New(domainErr.ErrOther)
	}

	ctx.SetCookie(
		"token",                    // name
		token,                      // value (JWT token string)
		int(duration.Seconds()),    // max age in seconds (e.g., 1 hour)
		"/",                        // path
		"",                         // domain ("" means current domain)
		c.Cfg.GinMode == "release", // secure (only send over HTTPS)
		true,                       // httpOnly (not accessible via JS)
	)

	return err
}

func (c *AuthController) Login(ctx *gin.Context) error {
	var input dto.LoginInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return err
	}

	token, err := c.AuthService.Login(input.Username, input.Password)
	if err != nil {
		return err
	}

	duration, err := time.ParseDuration(c.Cfg.JWTExpiresAt)
	if err != nil {
		return domainErr.New(domainErr.ErrOther)
	}

	ctx.SetCookie(
		"token",                    // name
		token,                      // value (JWT token string)
		int(duration.Seconds()),    // max age in seconds (e.g., 1 hour)
		"/",                        // path
		"",                         // domain ("" means current domain)
		c.Cfg.GinMode == "release", // secure (only send over HTTPS)
		true,                       // httpOnly (not accessible via JS)
	)

	return err

}
