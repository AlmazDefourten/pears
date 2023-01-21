package routing

import (
	"github.com/AlmazDefourten/goapp/models/container_models"
	"github.com/kataras/iris/v12"
)

type Router struct {
	HandlerContainer *container_models.HandlerContainer
}

func NewRouter(handlerContainer *container_models.HandlerContainer) *Router {
	return &Router{
		HandlerContainer: handlerContainer,
	}
}

// UseRoutes main API router
func (router *Router) UseRoutes(app *iris.Application) {
	userAPI := app.Party("/user")
	{
		userAPI.Use(iris.Compression)
		app.UseRouter(CorsHandler)
		app.Use(router.HandlerContainer.AuthHandler.AuthMiddleware)

		userAPI.Get("/list", router.HandlerContainer.UserInfoHandler.List)
		userAPI.Post("/registration", router.HandlerContainer.AuthHandler.Registration)
		userAPI.Post("/authorization", router.HandlerContainer.AuthHandler.Authorization)
	}
}

func InitializeRoutes(app *iris.Application, containerHandler container_models.HandlerContainer) {
	router := NewRouter(&containerHandler)

	router.UseRoutes(app)
}

// CorsHandler it`s middleware that handling requests: return nothing if it`s options
//	preflight request or go to handler if it`s other request type (get, post...)
func CorsHandler(ctx iris.Context) {
	if origin := ctx.GetHeader("Origin"); origin != "" {
		ctx.Header("Access-Control-Allow-Origin", origin)
		ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		ctx.Header("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
	// Stop here if its Preflighted OPTIONS request
	if ctx.Request().Method == "OPTIONS" {
		return
	}
	ctx.Next()
}
