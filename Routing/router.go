package routing

import (
	"github.com/AlmazDefourten/goapp/handlers"

	"github.com/kataras/iris/v12"
)

// There are all api routes for app
func UseRoutes(app *iris.Application) {
	userAPI := app.Party("/user")
	{
		userAPI.Use(iris.Compression)

		userAPI.Get("/list", handlers.List)
	}
}
