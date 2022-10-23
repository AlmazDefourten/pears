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

	migration.RunBaseMigration(container.AppConnection)

	routing.UseRoutes(app)

	app.Listen(":8080")
}
