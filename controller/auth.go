package controller

import (
	services "main/service"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService *services.AuthService
}

type RegisterInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{AuthService: authService}
}

func (c *AuthController) Register(ctx *gin.Context) error {
	var input RegisterInput
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
	var input LoginInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return err
	}

	err := c.AuthService.Login(input.Username, input.Password, ctx)

	return err

}
