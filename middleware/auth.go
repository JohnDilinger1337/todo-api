package middleware

import (
	"main/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware(jwtService *service.JWTService) gin.HandlerFunc {
    return func(c *gin.Context) {
      tokenString, _ := c.Cookie("token");
			 if tokenString == "" {
        c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
      }

			token, err := jwtService.ValidateToken(tokenString)
			 if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}
			
			claims := token.Claims.(jwt.MapClaims)
			c.Set("claims", claims)
			c.Next()
    }
}