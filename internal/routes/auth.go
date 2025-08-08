package routes

import (
	"github.com/gin-gonic/gin"
)

func authRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/sign-up")
		auth.POST("/login")
	}
}