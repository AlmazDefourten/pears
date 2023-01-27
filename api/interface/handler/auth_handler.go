package handler

import (
	"fmt"
	"github.com/AlmazDefourten/goapp/models"
	"github.com/kataras/iris/v12"
)

type AuthHandler struct {
	AuthService models.IAuthService
}

func NewAuthHandler(authService models.IAuthService) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
	}
}

func (authHandler *AuthHandler) Registration(ctx iris.Context) {
	var user models.User
	err := ctx.ReadJSON(&user)
	if err != nil {
		println(err)
		// logging here
	}
	ok := authHandler.AuthService.Registration(&user)
	response := map[string]interface{}{"result": ok}
	err = ctx.JSON(response)
	if err != nil {
		println(err)
		// logging here
	}
}

func (authHandler *AuthHandler) Authorization(ctx iris.Context) {
	var user models.User
	err := ctx.ReadJSON(&user)
	if err != nil {
		println(err)
		// logging here
	}
	ok, token := authHandler.AuthService.Authorization(user.Login, user.Password)
	response := map[string]interface{}{"result": ok, "token": token}
	err = ctx.JSON(response)
	if err != nil {
		println(err)
		// logging here
	}
}

func (authHandler *AuthHandler) AuthMiddleware(ctx iris.Context) {
	token := ctx.GetHeader("Token")
	ok, username := authHandler.AuthService.AuthCheck(token)
	if ok == false {
		err := ctx.JSON(models.Response{Ok: false, Message: "Пользователь не авторизован"})
		if err != nil {
			// logging here lol
			fmt.Println(username)
		}
	} else {
		ctx.Next()
	}
}
