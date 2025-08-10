package handlers

import "github.com/ThuraMinThein/golang-gin/internal/models"

type userHandler interface {
	Create(*models.User) (models.User, error)

}

type Handler struct {
	User userHandler
}