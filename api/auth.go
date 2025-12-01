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

func NewAuthAPI(authController controller.AuthController) *AuthAPI {
	return &AuthAPI{authController: &authController}
}

// Path information for swagger documentation
// @Summary Register a new user
// @Description Register a new user with username, email, and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body controller.RegisterInput true "User registration info"
// @Success 201 {object} controller.UserResponse
// @Failure 400 {object} controller.ErrorResponse
func (api *AuthAPI) RegisterRoute(ctx *gin.Context) {
	token, err := api.authController.Register(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"token": token})
}

func (api *AuthAPI) LoginRoute(ctx *gin.Context) {
	token, err := api.authController.Login(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
