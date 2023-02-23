package handler

import (
	"fmt"
	"github.com/AlmazDefourten/goapp/models"
	"github.com/kataras/iris/v12"
	"net/http"
)

type AuthHandler struct {
	AuthService models.IAuthService
}

func NewAuthHandler(authService models.IAuthService) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
	}
}

// Registration ShowAccount godoc
//
//	@Summary		Registration
//	@Description	add new user to the database
//	@Accept			json
//	@Produce		json
//	@Param			body		body		models.User		true	"request body with info about user"
//	@Failure		401	{object}	models.Response
//	@Success		200	{object}	models.Response
//	@Router			/user/registration [post]
//  swagger: example
func (authHandler *AuthHandler) Registration(ctx iris.Context) {
	var user models.User
	err := ctx.ReadJSON(&user)
	if err != nil {
		println(err)
		// logging here
	}
	ok := authHandler.AuthService.Registration(&user)
	response := models.Response{Ok: ok, Message: ""}
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
	if !ok {
		ctx.StatusCode(http.StatusUnauthorized)
	}
	response := map[string]interface{}{"ok": ok, "token": token}
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
		ctx.StopWithStatus(http.StatusUnauthorized)
		err := ctx.JSON(models.Response{Ok: false, Message: "Пользователь не авторизован"})
		if err != nil {
			// logging here lol
			fmt.Println(username)
		}
	} else {
		ctx.Next()
	}
}
