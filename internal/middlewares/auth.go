package middlewares

import (
	"net/http"
	"strings"

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
		
		if jwt != "abc" {
			unauthorizeError(c)
			return
		}
		
		c.Next()
        
    }
}

func unauthorizeError(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
    c.Abort()
}