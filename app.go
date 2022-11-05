package main

import (
	"fmt"
	iris "github.com/kataras/iris/v12"

	migration "github.com/AlmazDefourten/goapp/migrations"
	routing "github.com/AlmazDefourten/goapp/routing"
)

func main() {
	app := iris.New()

	// initialize global dependencies
	container := InitializeContainer()
	containerService := InitServiceDependency(&container)

	fmt.Println(containerService.UserService.List) // TODO: debug, delete this and create using of container

	migration.RunBaseMigration(container.AppConnection)

	routing.UseRoutes(app)

	app.Listen(":8080")
}
