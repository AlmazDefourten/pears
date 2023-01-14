//go:build wireinject
// +build wireinject

package container

import (
	"github.com/AlmazDefourten/goapp/infrastructure/configurator"
	"github.com/AlmazDefourten/goapp/interface/handler"
	"github.com/AlmazDefourten/goapp/models"
	"github.com/AlmazDefourten/goapp/services"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

// Initialize container with global app dependencies -
// Connection, configurator, etc...
func InitializeContainer() models.Container {
	wire.Build(NewContainer, NewConnection, configurator.NewViperConfigurator, wire.Bind(new(models.Configurator), new(*viper.Viper)))
	return models.Container{}
}

// Initialize dependencies for services
func InitServiceDependency(container_inited *models.Container) models.ServiceContainer {
	wire.Build(NewServiceContainer, services.NewUserService, wire.Bind(new(models.IUserService), new(*services.UserService)))
	return models.ServiceContainer{}
}

// RegisterServices - decomposition ServiceContainer to services
func RegisterServices(serviceContainer *models.ServiceContainer) HandlerContainer {
	return InitHandlerDependency(
		serviceContainer.UserService,
		serviceContainer.JWTService,
	)
}

// Initialize dependencies for handlers
func InitHandlerDependency(userService models.IUserService) HandlerContainer {
	wire.Build(NewHandlerContainer, handler.NewUserInfoHandler)
	return HandlerContainer{}
}
