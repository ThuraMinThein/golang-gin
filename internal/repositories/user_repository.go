package repositories

import (
	"github.com/ThuraMinThein/golang-gin/internal/database"
	"github.com/ThuraMinThein/golang-gin/internal/models"
)

func GetAllUser() ([]*models.UserWithToken, error) {
	var users []*models.UserWithToken
	err := database.DB.Find(&users).Error
	return users, err
}

func CreateUser(user *models.UserWithToken) error {
	return database.DB.Create(&user).Error
}

func GetOneUser(id uint) (*models.UserWithToken, error) {
	var user *models.UserWithToken
	err :=database.DB.First(&user, "id = ?", id).Error
	return user, err
}

func GetUserByUsername(username string) (*models.UserWithToken, error) {
	var user *models.UserWithToken
	err := database.DB.Find(&user, "user_name = ?", username).Error
	return user, err
}

func UpdateUser(user *models.UserWithToken) error {
	return database.DB.Save(user).Error
}