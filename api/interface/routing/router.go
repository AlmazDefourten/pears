package routing

import (
	_ "github.com/AlmazDefourten/goapp/docs"
	"github.com/AlmazDefourten/goapp/infrastructure/configurator"
	"github.com/AlmazDefourten/goapp/models/container_models"
	"github.com/iris-contrib/swagger"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"
)

const (
	apiPath = "/api/v1"
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
	app.UseRouter(CorsHandler)

	AutoDocHandleInit(app)

	userAPI := app.Party(apiPath + "/user")
	{
		userAPI.Use(iris.Compression)
		userAPI.Post("/registration", router.HandlerContainer.AuthHandler.Registration)
		userAPI.Post("/authorization", router.HandlerContainer.AuthHandler.Authorization)
	}
	userInfoAPI := app.Party("/userinfo")
	{
		userInfoAPI.UseRouter(router.HandlerContainer.AuthHandler.AuthMiddleware)
		userInfoAPI.Get("/list", router.HandlerContainer.UserInfoHandler.List)
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

// AutoDocHandleInit is routing and initializing autodocs
func AutoDocHandleInit(app *iris.Application) {
	// Configure the swagger UI page.
	config := configurator.SwaggerConfig
	swaggerUI := swagger.Handler(swaggerFiles.Handler, config)

	// Register on domain:port/swagger
	app.Get("/swagger", swaggerUI)
	// And the wildcard one for index.html, *.js, *.css and e.t.c.
	app.Get("/swagger/{any:path}", swaggerUI)
}
