package handlers

import (
	"github.com/kataras/iris/v12"
)

// Endpoint for a List of Users
func List(ctx iris.Context) {
	list := 5 // TODO: delete, hardcode
	ctx.JSON(list)
}
