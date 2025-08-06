package routes

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine) {

	user := r.Group("/users")
	{
		user.GET("/")
		user.POST("/")
	}

}