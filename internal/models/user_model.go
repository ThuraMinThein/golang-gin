package models

import (
	"github.com/ThuraMinThein/golang-gin/config/enums"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name			string				`json:"name"  binding:"required"`
	Email			string				`json:"email"  binding:"required" gorm:"uniqueIndex"`
	Password		string 				`json:"-"`
	UserName 		string				`json:"user_name"  binding:"required" gorm:"uniqueIndex"`
	Role 			enums.UserRoleEnum	`json:"role" binding:"required"`
	AccessToken 	string				`json:"access_token" gorm:"-"`
	RefreshToken 	string				`json:"refresh_token"`
}