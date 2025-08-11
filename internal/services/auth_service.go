package services

import (
	"github.com/ThuraMinThein/golang-gin/internal/dto"
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

	user.AccessToken, user.RefreshToken, err = helper.GetTokens(user.ID, &user.Role)

	if err != nil {
		return nil, err
	}

	err = repositories.CreateUser(user)
	
	if err != nil {
		return nil, err
	}

	return user,nil
}

func SignIn(loginDto *dto.LoginUserDto) (*models.User, error) {

	user, err := repositories.GetUserByUsername(loginDto.Username)

	if err != nil {
		return nil, err
	}

	err = helper.VerifyHashed(loginDto.Password, user.Password)

	if err != nil {
		return nil, err
	}

	user.AccessToken, user.RefreshToken, err = helper.GetTokens(user.ID, &user.Role)

	if err != nil {
		return nil, err
	}

	return user, nil
}