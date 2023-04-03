package handler

import (
	"fmt"
	"github.com/AlmazDefourten/goapp/infrastructure/loggerinstance"
	"github.com/AlmazDefourten/goapp/models"
	"github.com/golobby/container/v3"
	"github.com/kataras/iris/v12"
	"net/http"
)

type AuthHandler struct {
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
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
func (authHandler *AuthHandler) Registration(ctx iris.Context) {
	var user models.User
	err := ctx.ReadJSON(&user)
	if err != nil {
		loggerinstance.GlobalLogger.Error(err)
	}

	var authService models.IAuthService
	err = container.Resolve(&authService)
	if err != nil {
		loggerinstance.GlobalLogger.Error(err)
		panic(err)
	}

	ok, err := authService.Registration(&user)
	if err != nil {
		//loggine here
	}
	response := models.Response{Ok: ok, Message: ""}
	err = ctx.JSON(response)
	if err != nil {
		println(err)
		loggerinstance.GlobalLogger.Error(err)
	}
}

// Authorization ShowAccount godoc
//
//	@Summary		Authorization
//	@Description	authorization and take a token
//	@Accept			json
//	@Produce		json
//	@Param			body		body		models.UserAuthInfo		true	"request body with login and password"
//	@Failure		401	{object}	models.AuthResponse
//	@Success		200	{object}	models.AuthResponse
//	@Router			/user/authorization [post]
func (authHandler *AuthHandler) Authorization(ctx iris.Context) {
	var user models.User
	err := ctx.ReadJSON(&user)
	if err != nil {
		println(err)
		loggerinstance.GlobalLogger.Error(err)
	}

	var authService models.IAuthService
	err = container.Resolve(&authService)
	if err != nil {
		loggerinstance.GlobalLogger.Error(err)
		panic(err)
	}

	ok, token := authService.Authorization(user.Login, user.Password)
	if !ok {
		ctx.StatusCode(http.StatusUnauthorized)
	}
	var responseMessage string
	if ok {
		responseMessage = "Вы успешно авторизовались"
	} else {
		responseMessage = "Неверный логин или пароль"
	}
	response := models.AuthResponse{Ok: ok, Message: responseMessage, Access: token.AccessToken}

	ctx.SetCookie(&iris.Cookie{
		Name:     "refreshtoken",
		Value:    token.RefreshToken,
		HttpOnly: true,
	}, iris.CookieAllowSubdomains())
	err = ctx.JSON(response)
	if err != nil {
		println(err)
		loggerinstance.GlobalLogger.Error(err)
	}
}

// AuthMiddleware it`s middleware for check if user is authorized
func (authHandler *AuthHandler) AuthMiddleware(ctx iris.Context) {
	token := ctx.GetHeader("access_token")

	var authService models.IAuthService
	err := container.Resolve(&authService)
	if err != nil {
		loggerinstance.GlobalLogger.Error(err)
		panic(err)
	}

	ok, username := authService.AuthCheck(token)
	if ok == false {
		ctx.StopWithStatus(http.StatusUnauthorized)
		err := ctx.JSON(models.Response{Ok: false, Message: "Пользователь не авторизован"})
		if err != nil {
			loggerinstance.GlobalLogger.Error(err)
			fmt.Println(username)
		}
	} else {
		ctx.Next()
	}
}

// RefreshTokens it`s method for check refresh token and refresh tokens
func (authHandler *AuthHandler) RefreshTokens(ctx iris.Context) {
	token := ctx.GetHeader("refresh_token")

	var authService models.IAuthService
	err := container.Resolve(&authService)
	if err != nil {
		loggerinstance.GlobalLogger.Error(err)
		panic(err)
	}

	ok, tokens := authService.RefreshCheck(token)
	if ok == false {
		ctx.StopWithStatus(http.StatusUnauthorized)
		err := ctx.JSON(models.Response{Ok: false, Message: "Пользователь не авторизован"})
		if err != nil {
			loggerinstance.GlobalLogger.Error(err)
			fmt.Println(err)
		}
	} else {
		response := models.AuthResponse{Ok: ok, Access: tokens.AccessToken}

		ctx.SetCookie(&iris.Cookie{
			Name:     "refreshtoken",
			Value:    tokens.RefreshToken,
			HttpOnly: true,
		}, iris.CookieAllowSubdomains())

		err = ctx.JSON(response)
		if err != nil {
			println(err)
			loggerinstance.GlobalLogger.Error(err)
		}
	}
}
