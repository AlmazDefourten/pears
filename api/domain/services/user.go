package services

import (
	"github.com/AlmazDefourten/goapp/infra/logger_instance"
	"github.com/AlmazDefourten/goapp/models/repo_models"
	models "github.com/AlmazDefourten/goapp/models/user_models"
	"github.com/golobby/container/v3"
)

// UserService struct of service that works with Users
type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (userService *UserService) List() []models.User {
	var userRepo repo_models.UserRepository
	err := container.Resolve(&userRepo)
	if err != nil {
		logger_instance.GlobalLogger.Error("Ошибка при получении репозитория UserRepository")
	}
	usersList := userRepo.List()
	return usersList
}
