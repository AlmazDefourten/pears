//go:build wireinject
// +build wireinject

package main

import (
	"github.com/AlmazDefourten/goapp/container"
	models "github.com/AlmazDefourten/goapp/models"
	services "github.com/AlmazDefourten/goapp/services"
	"github.com/google/wire"
)

// Initialize container with global app dependencies -
// Connection, configurator, etc...
func InitializeContainer() models.Container {
	wire.Build(container.NewContainer, container.NewConnection, container.NewViperConfigurator)
	return models.Container{}
}

// Initialize services examples
func InitializeServiceContainer() models.Services {
	wire.Build(
		wire.NewSet(
			services.NewUserService,
			wire.Bind(new(models.IUserService), new(services.UserService)),
			container.NewServices,
		),
	)
	return models.Services{}
}
