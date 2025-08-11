package handlers

import (
	"net/http"

	"github.com/ThuraMinThein/golang-gin/internal/dto"
	"github.com/ThuraMinThein/golang-gin/internal/models"
	"github.com/ThuraMinThein/golang-gin/internal/services"
	"github.com/gin-gonic/gin"
)


func SignUp(c *gin.Context) {
	var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userData, err := services.SignUp(&user); 
	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

	c.SetCookie("refreshToken",userData.RefreshToken, 7 * 24 * 60 * 60, "/", "localhost", false, true)

    c.JSON(http.StatusCreated, userData)
}

func Login(c *gin.Context) {
	var dto dto.LoginUserDto
	if err := c.ShouldBindJSON(&dto); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	loginData, err := services.SignIn(&dto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = services.UpdateUser(loginData)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("refreshToken",loginData.RefreshToken, 7 * 24 * 60 * 60, "/", "localhost", false, true)
	c.JSON(http.StatusOK, loginData)

}