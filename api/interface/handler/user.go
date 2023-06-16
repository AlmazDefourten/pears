package handler

import (
	"github.com/AlmazDefourten/goapp/infra/logger_instance"
	"github.com/AlmazDefourten/goapp/models"
	"github.com/golobby/container/v3"
	"github.com/kataras/iris/v12"
)

// UserInfoHandler - handler for handle requests with user/s info
type UserInfoHandler struct {
}

func NewUserInfoHandler() *UserInfoHandler {
	return &UserInfoHandler{}
}

// List ShowAccount godoc
//
//	@Summary		List of users
//	@Description	take list of all users
//	@Accept			json
//	@Produce		json
//	@Failure		401	{object}	models.Response
//	@Success		200	{object}	[]models.User
//	@Router			/userinfo/list [get]
//  @Security 		JWTToken
func (userInfoHandler *UserInfoHandler) List(ctx iris.Context) {
	var userService models.IUserService
	err := container.Resolve(&userService)
	if err != nil {
		errRet := ctx.JSON(models.Response{Ok: false, Message: "Произошла ошибка, попробуйте позднее"})
		if errRet != nil {
			logger_instance.GlobalLogger.Error(err)
			return
		}
		logger_instance.GlobalLogger.Error(err)
	}
	data := userService.List()
	err = ctx.JSON(data)
	if err != nil {
		errRet := ctx.JSON(models.Response{Ok: false, Message: "Произошла ошибка, попробуйте позднее"})
		if errRet != nil {
			logger_instance.GlobalLogger.Error(errRet)
			return
		}
		logger_instance.GlobalLogger.Error(err)
	}
}
