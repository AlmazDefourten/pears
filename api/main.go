package main

import (
	"github.com/AlmazDefourten/goapp/infrastructure/container"
	"github.com/AlmazDefourten/goapp/infrastructure/migrations"
	"github.com/AlmazDefourten/goapp/interface/routing"
	"github.com/kataras/iris/v12"
)

// @title Pears auto documentation
// @version 1.0
// @description Pears API, specification and description

// @host localhost:8080
// @BasePath /api

// @securityDefinitions.apikey	JWTToken
// @in							header
// @name						token
// @description				Access token only
func main() {
	app := iris.New()

	initializeApp(app)
}

func initializeApp(app *iris.Application) {
	_container := container.InitializeContainer()
	containerService := container.InitServiceDependency(&_container)

	migrations.RunBaseMigration(_container.AppConnection)

	containerHandler := container.RegisterServices(containerService)
	routing.InitializeRoutes(app, containerHandler)
	err := app.Listen(":" + _container.ConfigProvider.GetString("host_port"))
	if err != nil {
		// there is logging
		panic(err)
	}
}
