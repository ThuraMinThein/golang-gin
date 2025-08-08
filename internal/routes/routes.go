package routes

import (
	"github.com/ThuraMinThein/golang-gin/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.Use(middlewares.ErrorHandler())

	authRoutes(r)
	userRoutes(r);

}