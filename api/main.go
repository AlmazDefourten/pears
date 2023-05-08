package main

import (
	"github.com/AlmazDefourten/goapp/infrastructure/logger_instance"
	"github.com/AlmazDefourten/goapp/infrastructure/migrations"
	"github.com/AlmazDefourten/goapp/infrastructure/resolver"
	"github.com/AlmazDefourten/goapp/interface/routing"
	"github.com/AlmazDefourten/goapp/models"
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
	err := resolver.InitializeContainer()
	if err != nil {
		panic(err)
	}
	err = resolver.RegisterServices()
	if err != nil {
		logger_instance.GlobalLogger.Error(err)
		panic(err)
	}
	migrations.RunBaseMigration()
	routing.InitializeRoutes(app)
	var c models.Configurator
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
