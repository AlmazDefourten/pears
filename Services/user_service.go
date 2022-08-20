package userservice

import (
	usermodel "github.com/AlmazDefourten/goapp/models"
)

func List() []usermodel.User {
	users := []usermodel.User{
		{"Meow", 18},
		{"Hi!", 20},
	}
	return users
}
