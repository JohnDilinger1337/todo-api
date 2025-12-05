package api

import (
	"main/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

var ErrorUnauthorized = gin.H{"message": "Something went wrong while authenticating user!"}
var ErrorForbidden = gin.H{"message": "forbidden"}
var ErrorBadRequest = gin.H{"message": "Something went wrong while processing your request!"}
var ErrorUserAlreadyLoggedIn = gin.H{"message": "user already logged in!"}

var ResponseLoggedIn = gin.H{"message": "Logged in successfully"}
var ResponseRegistered = gin.H{"message": "Registered successfully! You're now logged in."}

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
// @Success 201 {object} ResponseRegistered
// @Failure 400 {object} ErrorBadRequest
// @Router /auth/register [post]
func (api *AuthAPI) RegisterRoute(ctx *gin.Context) {
	err := api.authController.Register(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorBadRequest)
		return
	}
	ctx.JSON(http.StatusCreated, ResponseRegistered)
}

// LoginRoute godoc
// @Summary login a user
// @Description Login a user with username and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body controller.LoginInput true "User login info"
// @Success 200 {object} ResponseLoggedIn
// @Failure 400 {object} ErrorUnauthorized
// @Router /auth/login [post]
func (api *AuthAPI) LoginRoute(ctx *gin.Context) {
	cookie, _ := ctx.Cookie("token")
	if cookie != "" {
		ctx.JSON(http.StatusBadRequest, ErrorUserAlreadyLoggedIn)
		return
	}

	err := api.authController.Login(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, ErrorUnauthorized)
		return
	}
	// Print Set-Cookie header to verify cookie is set in response
	ctx.JSON(http.StatusOK, ResponseLoggedIn)
}
