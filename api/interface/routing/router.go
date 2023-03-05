package routing

import (
	_ "github.com/AlmazDefourten/goapp/docs"
	"github.com/AlmazDefourten/goapp/infrastructure/configurator"
	"github.com/AlmazDefourten/goapp/interface/handler"
	"github.com/iris-contrib/swagger"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"
)

const (
	apiPath = "/api"
)

type Router struct {
}

func NewRouter() *Router {
	return &Router{}
}

// UseRoutes main API router
func (router *Router) UseRoutes(app *iris.Application) {
	app.UseRouter(CorsHandler)

	AutoDocHandleInit(app)

	var authHandler = handler.NewAuthHandler()
	// TODO: refactor this by adding new party with apipath
	userAPI := app.Party(apiPath + "/user")
	{
		userAPI.Use(iris.Compression)

		userAPI.Post("/registration", authHandler.Registration)
		userAPI.Post("/authorization", authHandler.Authorization)
	}
	userInfoAPI := app.Party(apiPath + "/userinfo")
	{
		userInfoAPI.UseRouter(authHandler.AuthMiddleware)

		var userInfoHandler = handler.NewUserInfoHandler()

		userInfoAPI.Get("/list", userInfoHandler.List)
	}
	postAPI := app.Party(apiPath + "/post")
	{
		var postHandler = handler.NewPostHandler()
		postAPI.UseRouter(authHandler.AuthMiddleware)
		postAPI.Post("/create", postHandler.Create)
		postAPI.Get("/list", postHandler.List)
		postAPI.Get("/get", postHandler.Get)
	}
}

func InitializeRoutes(app *iris.Application) {
	router := NewRouter()

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
