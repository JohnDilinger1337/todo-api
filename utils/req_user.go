package utils

import (
	"main/middleware"

	"github.com/gin-gonic/gin"
)

func ReqUser(c *gin.Context) middleware.AuthUser {
	userAny, exists := c.Get("user")
	if !exists {
		return middleware.AuthUser{}
	}
	user := userAny.(middleware.AuthUser)
	return user
}
