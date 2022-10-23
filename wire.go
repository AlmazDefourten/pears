//go:build wireinject
// +build wireinject

package main

import (
	"github.com/AlmazDefourten/goapp/container"
	"github.com/google/wire"
)

// Initialize container with global app dependencies -
// Connection, configurator, etc...
func InitializeContainer() container.Container {
	wire.Build(container.NewContainer, container.NewConnection, container.NewViperConfigurator)
	return container.Container{}
}
