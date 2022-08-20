package handlers

import (
	userservice "github.com/AlmazDefourten/goapp/services"

	"github.com/kataras/iris/v12"
)

func List(ctx iris.Context) {
	list := userservice.List()
	ctx.JSON(list)
}
