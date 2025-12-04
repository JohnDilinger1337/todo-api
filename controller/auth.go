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

type RegisterResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{AuthService: authService}
}

func (c *AuthController) Register(ctx *gin.Context) (string, error) {
	var input RegisterInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return "", err
	}

	user, err := c.AuthService.Register(input.Username, input.Email, input.Password)

	if err != nil {
		return "", err
	}

	token, err := c.AuthService.Login(user.Username, user.Password)

	if err != nil {
		return "", err
	}

	return token, err
}

func (c *AuthController) Login(ctx *gin.Context) (string, error) {
	var input LoginInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return "", err
	}

	token, err := c.AuthService.Login(input.Username, input.Password)
	if err != nil {
		return "", err
	}

	return token, err
}
