package main

import (
	"github.com/AlmazDefourten/goapp/infrastructure/migrations"
	"github.com/AlmazDefourten/goapp/interface/routing"
	iris "github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	_container := InitializeContainer()
	containerService := InitServiceDependency(&_container)

	migrations.RunBaseMigration(_container.AppConnection)

	containerHandler := RegisterServices(&containerService)
	routing.InitializeRoutes(app, containerHandler)
	app.Listen(":8080")
}
