package routes

import (
	"github.com/ThuraMinThein/golang-gin/internal/handlers"
	"github.com/ThuraMinThein/golang-gin/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func userRoutes(r *gin.Engine) {
	user := r.Group("/users")
	user.Use(middlewares.AuthMiddleware())
	{
		user.GET("/", handlers.GetUsers)
		user.GET("/:id")
		user.PATCH("/:id")
		user.DELETE("/:id")
	}
}