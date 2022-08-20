package main

import (
	"github.com/kataras/iris/v12"

	"github.com/AlmazDefourten/goapp/routing"
)

func main() {
	app := iris.New()

	routing.UseRoutes(app)

	app.Listen(":8080")
}
