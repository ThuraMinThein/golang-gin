package handlers

import (
	"net/http"

	"github.com/ThuraMinThein/golang-gin/internal/services"
	"github.com/gin-gonic/gin"
)


func GetUsers(c *gin.Context) {
	users, err := services.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func GetOneUser(c *gin.Context, id uint) {
	user, err := services.GetOneUser(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}