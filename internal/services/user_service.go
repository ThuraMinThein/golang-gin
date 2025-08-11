package services

import (
	"github.com/ThuraMinThein/golang-gin/internal/models"
	"github.com/ThuraMinThein/golang-gin/internal/repositories"
)

func GetUsers() ([]*models.User, error) {
	return repositories.GetAllUser()
}

func GetOneUser(id uint) (*models.User, error) {
	return repositories.GetOneUser(id)
}

func UpdateUser(user *models.User) error {
	return repositories.UpdateUser(user)
}