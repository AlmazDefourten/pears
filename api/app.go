package main

import (
	iris "github.com/kataras/iris/v12"

	migration "github.com/AlmazDefourten/goapp/migrations"
	routing "github.com/AlmazDefourten/goapp/routing"
)

func main() {
	app := iris.New()

	// initialize global dependencies
	container := InitializeContainer()
	containerService := InitServiceDependency(&container)

	migration.RunBaseMigration(container.AppConnection)

	containerHandler := RegisterServices(&containerService)

	router := routing.NewRouter(&containerHandler)

	router.UseRoutes(app)

	app.Listen(":8080")
}
