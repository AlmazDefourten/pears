package main

import (
	"github.com/AlmazDefourten/goapp/infrastructure/container"
	"github.com/AlmazDefourten/goapp/infrastructure/migrations"
	"github.com/AlmazDefourten/goapp/interface/routing"
	iris "github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	_container := container.InitializeContainer()
	containerService := container.InitServiceDependency(&_container)

	migrations.RunBaseMigration(_container.AppConnection)

	containerHandler := container.RegisterServices(&containerService)
	routing.InitializeRoutes(app, containerHandler)
	err := app.Listen(":" + _container.ConfigProvider.GetString("host_port"))
	if err != nil {
		// there is logging
		panic(err)
	}
}
