package main

import (
	"github.com/AlmazDefourten/goapp/infrastructure/container"
	"github.com/AlmazDefourten/goapp/infrastructure/migrations"
	"github.com/AlmazDefourten/goapp/interface/routing"
	"github.com/AlmazDefourten/goapp/pkg/logging/resolvers"
	"github.com/kataras/iris/v12"
)

// @title Pears auto documentation
// @version 1.0
// @description Pears API, specification and description

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey	JWTToken
// @in							header
// @name						token
// @description				Access token only
func main() {
	app := iris.New()

	initializeApp(app)
}

func initializeApp(app *iris.Application) {
	_container := container.InitializeContainer(resolvers.GlobalLogger)
	containerService := container.InitServiceDependency(&_container)

	log := container.InitLogrusLogger(resolvers.GlobalLogger)
	log.Info("Start app!")
	migrations.RunBaseMigration(_container.AppConnection)

	containerHandler := container.RegisterServices(containerService)
	routing.InitializeRoutes(app, containerHandler)
	err := app.Listen(":" + _container.ConfigProvider.GetString("host_port"))
	if err != nil {
		// there is logging
		panic(err)
	}
}
