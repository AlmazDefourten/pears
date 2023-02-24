//go:build wireinject
// +build wireinject

package container

import (
	"github.com/AlmazDefourten/goapp/infrastructure/configurator"
	"github.com/AlmazDefourten/goapp/infrastructure/data_adapter/connection"
	"github.com/AlmazDefourten/goapp/interface/handler"
	"github.com/AlmazDefourten/goapp/models"
	"github.com/AlmazDefourten/goapp/models/container_models"
	"github.com/AlmazDefourten/goapp/pkg/logging/loggers"
	"github.com/AlmazDefourten/goapp/pkg/logging/resolvers"
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
	wire.Build(NewServiceContainer, services.NewUserService, wire.Bind(new(models.IUserService), new(*services.UserService)),
		services.NewJWTService, wire.Bind(new(models.IJWTService), new(*services.JWTService)),
		services.NewAuthService, wire.Bind(new(models.IAuthService), new(*services.AuthService)),
	)
	return container_models.ServiceContainer{}
}

// RegisterServices - decomposition ServiceContainer to services
func RegisterServices(serviceContainer container_models.ServiceContainer) container_models.HandlerContainer {
	return InitHandlerDependency(
		serviceContainer.UserService,
		serviceContainer.AuthService,
	)
}

// Initialize dependencies for handlers
func InitHandlerDependency(userService models.IUserService, authService models.IAuthService) container_models.HandlerContainer {
	wire.Build(NewHandlerContainer, handler.NewUserInfoHandler, handler.NewAuthHandler)
	return container_models.HandlerContainer{}
}

func InitLogrusLogger(typeLogger resolvers.TypeLogger) models.Logger {
	wire.Build(wire.Bind(new(models.Logger), loggers.InitLogrus(typeLogger)))
	return loggers.InitLogrus(typeLogger)
}
