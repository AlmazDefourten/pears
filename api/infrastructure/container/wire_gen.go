// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package container

import (
	"github.com/AlmazDefourten/goapp/infrastructure/configurator"
	"github.com/AlmazDefourten/goapp/interface/handler"
	"github.com/AlmazDefourten/goapp/models"
	"github.com/AlmazDefourten/goapp/models/container_models"
	"github.com/AlmazDefourten/goapp/pkg/connection"
	"github.com/AlmazDefourten/goapp/services"
)

// Injectors from wire.go:

// Initialize container with global app dependencies -
// Connection, configurator, etc...
func InitializeContainer() container_models.Container {
	viper := configurator.NewViperConfigurator()
	db := connection.NewGormConnection(viper)
	container := NewContainer(db, viper)
	return container
}

// Initialize dependencies for services
func InitServiceDependency(container_inited *container_models.Container) container_models.ServiceContainer {
	userService := services.NewUserService(container_inited)
	serviceContainer := NewServiceContainer(userService)
	return serviceContainer
}

// Initialize dependencies for handlers
func InitHandlerDependency(userService models.IUserService) container_models.HandlerContainer {
	userInfoHandler := handler.NewUserInfoHandler(userService)
	handlerContainer := NewHandlerContainer(userInfoHandler)
	return handlerContainer
}

// wire.go:

// RegisterServices - decomposition ServiceContainer to services
func RegisterServices(serviceContainer *container_models.ServiceContainer) container_models.HandlerContainer {
	return InitHandlerDependency(
		serviceContainer.UserService,
	)
}