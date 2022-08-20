package routing

import (
	"github.com/AlmazDefourten/goapp/handlers"

	"github.com/kataras/iris/v12"
)

func UseRoutes(app *iris.Application) {
	userAPI := app.Party("/user")
	{
		userAPI.Use(iris.Compression)

		userAPI.Get("/list", handlers.List)
	}
}
