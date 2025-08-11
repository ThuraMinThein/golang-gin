package services

import (
	"github.com/ThuraMinThein/golang-gin/internal/models"
	"github.com/ThuraMinThein/golang-gin/internal/repositories"
)

func GetUsers() ([]*models.UserWithToken, error) {
	return repositories.GetAllUser()
}

func GetOneUser(id uint) (*models.UserWithToken, error) {
	return repositories.GetOneUser(id)
}

func UpdateUser(id uint, user *models.UserWithToken) error {
	userData, err := repositories.GetOneUser(id)
	if err != nil {
		return err
	}
	return repositories.UpdateUser(userData)
}