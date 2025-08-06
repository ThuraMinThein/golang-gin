package services

import (
	"github.com/ThuraMinThein/golang-gin/internal/models"
	"github.com/ThuraMinThein/golang-gin/internal/repositories"
)

func GetUsers() ([]models.User, error) {
	return repositories.GetAllUser()
}

func CreateUser(user *models.User) error {
	return repositories.CreateUser(user)
}