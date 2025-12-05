package controller

import (
	"main/dto"
	"main/service"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{AuthService: authService}
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

	err = c.AuthService.Login(user.Username, user.Password, ctx)

	return err
}

func (c *AuthController) Login(ctx *gin.Context) error {
	var input dto.LoginInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return err
	}

	err := c.AuthService.Login(input.Username, input.Password, ctx)

	return err

}
