package userservice

import (
	usermodel "github.com/AlmazDefourten/goapp/models"
)

// Service for operations with Users
type UserService interface {
}

// Get a List of Users
func List() []usermodel.User {
	users := []usermodel.User{
		{Name: "Meow", Age: 18},
		{Name: "Hi!", Age: 20},
	}
	return users
}
