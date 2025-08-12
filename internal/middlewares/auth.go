package middlewares

import (
	"net/http"
	"strings"

	"github.com/ThuraMinThein/golang-gin/internal/helper"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	 return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
		tokenType := strings.Split(token, " ")
		
		if token == "" || len(tokenType) != 2 || tokenType[0] != "Bearer" {
			unauthorizeError(c)
			return
		}

		jwt := tokenType[1]

		claims, err := helper.ParseToken(jwt)
		
		if err != nil {
			unauthorizeError(c)
			return
		}
		
		c.Set("user", claims)

		
		c.Next()
        
    }
}

func unauthorizeError(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
    c.Abort()
}