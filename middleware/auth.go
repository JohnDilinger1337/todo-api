package middleware

import (
	"main/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthUser struct {
	ID      uint `json:"id"`
	IsAdmin bool `json:"is_admin"`
}

func JWTMiddleware(jwtService *service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")
		if err != nil || tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			return
		}

		token, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Cannot parse token"})
			return
		}

		authUser := AuthUser{
			ID:      uint(claims["user_id"].(float64)),
			IsAdmin: claims["is_admin"].(bool),
		}

		c.Set("authUser", authUser)
		c.Next()
	}
}
