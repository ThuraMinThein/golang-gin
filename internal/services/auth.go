package services

import (
	"github.com/ThuraMinThein/golang-gin/internal/helper"
	"github.com/ThuraMinThein/golang-gin/internal/models"
	"github.com/ThuraMinThein/golang-gin/internal/repositories"
)

func SignUp(user *models.User) (*models.User, error) {

	hashedPassword, err := helper.Hash(user.Password)

	if err != nil {
		return nil, err
	}

	user.Password = hashedPassword

	err = repositories.CreateUser(user)
	
	if err != nil {
		return nil, err
	}

	user.AccessToken, user.RefreshToken, err = helper.GetTokens(user.ID, &user.Role)

	if err != nil {
		return nil, err
	}

	return user,nil
}

// func SignIn(username, password string) (string, error) {

// }