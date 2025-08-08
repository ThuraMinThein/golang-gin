package repositories

import (
	"github.com/ThuraMinThein/golang-gin/internal/database"
	"github.com/ThuraMinThein/golang-gin/internal/models"
)

func GetAllUser() ([]models.User, error) {
	var users []models.User
	err := database.DB.Find(&users).Error
	return users, err
}

func CreateUser(user *models.User) error {
	return database.DB.Create(user).Error
}

func GetOneUser(id uint) (models.User, error) {
	var user models.User
	err :=database.DB.Find(&user, "ID = ?", id).Error
	return user, err
}