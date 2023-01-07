//go:build wireinject
// +build wireinject

package container

import (
	"github.com/AlmazDefourten/goapp/infrastructure/configurator"
	"github.com/AlmazDefourten/goapp/interface/handler"
	"github.com/AlmazDefourten/goapp/models"
	"github.com/AlmazDefourten/goapp/models/container_models"
	"github.com/AlmazDefourten/goapp/pkg/connection"
	"github.com/AlmazDefourten/goapp/services"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

// Initialize container with global app dependencies -
// Connection, configurator, etc...
func InitializeContainer() container_models.Container {
	wire.Build(NewContainer, connection.NewGormConnection, configurator.NewViperConfigurator, wire.Bind(new(models.Configurator), new(*viper.Viper)))
	return container_models.Container{}
}

// Initialize dependencies for services
func InitServiceDependency(container_inited *container_models.Container) container_models.ServiceContainer {
	wire.Build(NewServiceContainer, services.NewUserService, wire.Bind(new(models.IUserService), new(*services.UserService)))
	return container_models.ServiceContainer{}
}

// RegisterServices - decomposition ServiceContainer to services
func RegisterServices(serviceContainer *container_models.ServiceContainer) container_models.HandlerContainer {
	return InitHandlerDependency(
		serviceContainer.UserService,
	)
}

// Initialize dependencies for handlers
func InitHandlerDependency(userService models.IUserService) container_models.HandlerContainer {
	wire.Build(NewHandlerContainer, handler.NewUserInfoHandler)
	return container_models.HandlerContainer{}
}
