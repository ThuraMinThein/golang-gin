package dto

type LoginUserDto struct {
	Username	string	`json:"user_name" bind:"required"`
	Password 	string	`json:"password" bind:"required"`
}