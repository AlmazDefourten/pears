package services

import (
	models "github.com/AlmazDefourten/goapp/models"
)

// UserService struct of service that works with Users
type UserService struct {
	Container *models.Container
}

func NewUserService(container *models.Container) *UserService {
	return &UserService{
		Container: container,
	}
}

// List Get a List of Users
func (userService *UserService) List() []models.User {
	var users []models.User
	userService.Container.AppConnection.Find(&users)
	return users
}
