//go:build wireinject
// +build wireinject

package main

import (
	"github.com/AlmazDefourten/goapp/container"
	"github.com/AlmazDefourten/goapp/handlers"
	"github.com/AlmazDefourten/goapp/models"
	"github.com/AlmazDefourten/goapp/services"
	"github.com/google/wire"
)

// Initialize container with global app dependencies -
// Connection, configurator, etc...
func InitializeContainer() models.Container {
	wire.Build(container.NewContainer, container.NewConnection, container.NewViperConfigurator)
	return models.Container{}
}

// Initialize dependencies for services
func InitServiceDependency(container_inited *models.Container) models.ServiceContainer {
	wire.Build(container.NewServiceContainer, services.NewUserService, wire.Bind(new(models.IUserService), new(*services.UserService)))
	return models.ServiceContainer{}
}

// RegisterServices - decomposition ServiceContainer to services
func RegisterServices(serviceContainer *models.ServiceContainer) container.HandlerContainer {
	return InitHandlerDependency(
		serviceContainer.UserService,
	)
}

// Initialize dependencies for handlers
func InitHandlerDependency(userService models.IUserService) container.HandlerContainer {
	wire.Build(container.NewHandlerContainer, handlers.NewUserInfoHandler)
	return container.HandlerContainer{}
}
