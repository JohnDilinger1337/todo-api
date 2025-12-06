package api

import (
	"main/controller"
	domainErr "main/domain/error"
	"main/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
// @Param user body dto.RegisterInput true "User registration info"
// @Success 201 {object} dto.SuccessMessageResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /auth/register [post]
func (api *AuthAPI) RegisterRoute(ctx *gin.Context) {
	err := api.authController.Register(ctx)

	if err != nil {
		if e, ok := err.(*domainErr.DomainError); ok {
			switch e.Code {
			case domainErr.ErrUserExistsCode:
				ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: dto.MsgUserExists})
				return
			case domainErr.ErrOther:
				ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: dto.MsgBadRequest})
				return
			}
		}
	}
	ctx.JSON(http.StatusCreated, dto.SuccessMessageResponse{Message: dto.MsgRegistered})
}

// LoginRoute godoc
// @Summary login a user
// @Description Login a user with username and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body dto.LoginInput true "User login info"
// @Success 200 {object} dto.SuccessMessageResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /auth/login [post]
func (api *AuthAPI) LoginRoute(ctx *gin.Context) {
	cookie, _ := ctx.Cookie("token")
	if cookie != "" {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: dto.MsgUserAlreadyLoggedIn})
		return
	}

	err := api.authController.Login(ctx)
	if err != nil {
		if e, ok := err.(*domainErr.DomainError); ok {
			switch e.Code {
			case domainErr.ErrUserNotFoundCode:
				ctx.JSON(http.StatusUnauthorized, dto.ErrorResponse{Message: dto.MsgUnauthorized})
				return
			case domainErr.ErrInvalidPasswordCode:
				ctx.JSON(http.StatusUnauthorized, dto.ErrorResponse{Message: dto.MsgInvalidPassword})
				return
			}
		}
		return
	}
	// Print Set-Cookie header to verify cookie is set in response
	ctx.JSON(http.StatusOK, dto.SuccessMessageResponse{Message: dto.MsgLoggedIn})
}
