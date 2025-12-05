package utils

import (
	"errors"
	"main/middleware"

	"github.com/gin-gonic/gin"
)

func ReqUser(c *gin.Context) (middleware.AuthUser, error) {
	userAny, exists := c.Get("user")
	if !exists {
		return middleware.AuthUser{}, errors.New("user not found")
	}

	user, ok := userAny.(middleware.AuthUser)
	if !ok {
		return middleware.AuthUser{}, errors.New("invalid user type")
	}
	return user, nil
}
