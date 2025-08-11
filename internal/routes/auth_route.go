package routes

import (
	"github.com/ThuraMinThein/golang-gin/internal/handlers"
	"github.com/gin-gonic/gin"
)

func authRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/sign-up", handlers.SignUp)
		auth.POST("/login", handlers.Login)
		auth.POST("/refresh")
	}
}