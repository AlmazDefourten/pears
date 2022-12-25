package routing

import (
	"github.com/AlmazDefourten/goapp/infrastructure/container"
	"github.com/kataras/iris/v12"
)

type Router struct {
	HandlerContainer *container.HandlerContainer
}

func NewRouter(handlerContainer *container.HandlerContainer) *Router {
	return &Router{
		HandlerContainer: handlerContainer,
	}
}

// UseRoutes main API router
func (router *Router) UseRoutes(app *iris.Application) {
	userAPI := app.Party("/user")
	{
		userAPI.Use(iris.Compression)

		userAPI.Get("/list", router.HandlerContainer.UserInfoHandler.List)
	}
}

func InitializeRoutes(app *iris.Application, containerHandler container.HandlerContainer) {
	router := NewRouter(&containerHandler)

	router.UseRoutes(app)
}
