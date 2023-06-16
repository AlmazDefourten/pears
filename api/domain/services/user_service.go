package services

import (
	models "github.com/AlmazDefourten/goapp/models/user_models"
	"github.com/golobby/container/v3"
	"gorm.io/gorm"
)

// UserService struct of service that works with Users
type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (userService *UserService) List() []models.User {
	var db gorm.DB
	err := container.Resolve(&db)
	if err != nil {
		return nil
	}
	var users []models.User
	db.Find(&users)
	return users
}
