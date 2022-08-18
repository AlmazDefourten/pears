package routing

import (
	"github.com/kataras/iris/v12"
)

func use_routes(app *iris.Application) {
	userAPI := app.Party("/user")
	{
		userAPI.Use(iris.Compression)

		userAPI.Get("/list", list)
	}
}
