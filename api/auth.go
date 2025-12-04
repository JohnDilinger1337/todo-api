package api

import (
	"main/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

var ErrorUnauthorized = gin.H{"error": "unauthorized"}
var ErrorForbidden = gin.H{"error": "forbidden"}
var ErrorBadRequest = gin.H{"error": "bad request"}

type AuthAPI struct {
	authController *controller.AuthController
}

func NewAuthAPI(authController *controller.AuthController) *AuthAPI {
	return &AuthAPI{authController: authController}
}

// RegisterRoute godoc
// @Summary Register a new user
// @Description Register a new user with username, email, and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body controller.RegisterInput true "User registration info"
// @Success 201 {object} controller.RegisterResponse
// @Failure 400 {object} controller.ErrorResponse
// @Router /auth/register [post]
func (api *AuthAPI) RegisterRoute(ctx *gin.Context) {
	token, err := api.authController.Register(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, controller.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"token": token})
}

// LoginRoute godoc
// @Summary login a user
// @Description Login a user with username and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body controller.LoginInput true "User login info"
// @Success 200 {object} controller.LoginResponse
// @Failure 400 {object} controller.ErrorResponse
// @Router /auth/login [post]
func (api *AuthAPI) LoginRoute(ctx *gin.Context) {
	token, err := api.authController.Login(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, controller.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
