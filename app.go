package main

import (
	migration "github.com/AlmazDefourten/goapp/migrations"
	"github.com/AlmazDefourten/goapp/pkg/logging"
	routing "github.com/AlmazDefourten/goapp/routing"
	iris "github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	logger := logging.GetLogger()
	logger.Info("start")
	// initialize global dependencies
	container := InitializeContainer()
	containerService := InitServiceDependency(&container)

	migration.RunBaseMigration(container.AppConnection)

	containerHandler := RegisterServices(&containerService)

	router := routing.NewRouter(&containerHandler)

	router.UseRoutes(app)

	app.Listen(":8080")
}
