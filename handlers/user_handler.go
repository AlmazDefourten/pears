package handlers

import (
	services "github.com/AlmazDefourten/goapp/services"

	"github.com/kataras/iris/v12"
)

// Endpoint for a List of Users
func List(ctx iris.Context) {
	list := services.List()
	ctx.JSON(list)
}
