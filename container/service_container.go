package container

import (
	models "github.com/AlmazDefourten/goapp/models"
)

func NewServices(userService models.IUserService) models.Services {
	return models.Services{
		UserService: userService,
	}
}
