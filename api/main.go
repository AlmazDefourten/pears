package main

import (
	"github.com/AlmazDefourten/goapp/infra/dependencies_registration"
	"github.com/AlmazDefourten/goapp/infra/logger_instance"
	"github.com/AlmazDefourten/goapp/infra/migrations"
	"github.com/AlmazDefourten/goapp/interface/routing"
	"github.com/AlmazDefourten/goapp/models/util_adapters"
	"github.com/golobby/container/v3"
	"github.com/kataras/iris/v12"
)

// @title Pears auto documentation
// @version 1.0
// @description Pears API, specification and description

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
	err := dependencies_registration.InitializeContainer()
	if err != nil {
		panic(err)
	}
	err = dependencies_registration.RegisterServices()
	if err != nil {
		logger_instance.GlobalLogger.Error(err)
		panic(err)
	}
	migrations.RunBaseMigration()
	routing.InitializeRoutes(app)
	var c util_adapters.Configurator
	err = container.Resolve(&c)
	if err != nil {
		logger_instance.GlobalLogger.Error(err)
		panic(err)
	}
	err = app.Listen(c.GetString("appserver") + ":" + c.GetString("host_port"))
	if err != nil {
		logger_instance.GlobalLogger.Error(err)
		panic(err)
	}
}
