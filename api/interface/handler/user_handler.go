package handler

import (
	"github.com/AlmazDefourten/goapp/models"
	"github.com/kataras/iris/v12"
)

// UserInfoHandler - handler for handle requests with user/s info
type UserInfoHandler struct {
	UserService models.IUserService
}

func NewUserInfoHandler(userService models.IUserService) *UserInfoHandler {
	return &UserInfoHandler{
		UserService: userService,
	}
}

// List Endpoint for a List of Users
func (userInfoHandler *UserInfoHandler) List(ctx iris.Context) {
	data := userInfoHandler.UserService.List()
	err := ctx.JSON(data)
	if err != nil {
		// there is logging
		panic(err)
	}
}
