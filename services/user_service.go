package services

import (
	models "github.com/AlmazDefourten/goapp/models"
)

type UserService struct {
	Container *models.Container
}

func NewUserService(container *models.Container) UserService {
	return UserService{
		Container: container,
	}
}

// Get a List of Users
func (service UserService) List() []models.User {
	var users []models.User
	result := service.Container.AppConnection.Find(&users)

	rows := result.RowsAffected
	rows++
	return []models.User{}
}
