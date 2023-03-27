package resolver

import (
	"github.com/AlmazDefourten/goapp/infrastructure/configurator"
	"github.com/AlmazDefourten/goapp/infrastructure/data_adapter/connection"
	"github.com/AlmazDefourten/goapp/models"
	"github.com/AlmazDefourten/goapp/pkg/logging/loggers"
	"github.com/AlmazDefourten/goapp/pkg/logging/resolvers"
	"github.com/AlmazDefourten/goapp/services"
	"github.com/golobby/container/v3"
	"gorm.io/gorm"
)

var globalLogger models.Logger

// InitializeContainer with global app dependencies -
// Connection, configurator, etc...
func InitializeContainer() error {
	err := container.NamedSingleton("globalLogger", func() models.Logger {
		return loggers.Init(resolvers.GlobalLogger)
	})
	err = container.NamedResolve(&globalLogger, "globalLogger")
	if err != nil {
		globalLogger.Error(err)
		return err
	}
	err = container.NamedSingleton("serviceLogger", func() models.Logger {
		return loggers.Init(resolvers.ServiceLogger)
	})
	if err != nil {
		globalLogger.Error(err)
		return err
	}

	err = container.Singleton(func() models.Configurator {
		return configurator.NewViperConfigurator()
	})
	if err != nil {
		globalLogger.Error(err)
		return err
	}
	var c models.Configurator
	err = container.Resolve(&c)
	if err != nil {
		globalLogger.Error(err)
		return err
	}

	err = container.Singleton(func() gorm.DB {
		return *connection.NewGormConnection(c)
	})
	if err != nil {
		globalLogger.Error(err)
		return err
	}
	return nil
}

// RegisterServices - decomposition ServiceContainer to services
func RegisterServices() error {
	err := container.Singleton(func() models.IUserService {
		return services.NewUserService()
	})
	if err != nil {
		globalLogger.Error(err)
		return err
	}

	err = container.Singleton(func() models.IJWTService {
		return services.NewJWTService()
	})
	if err != nil {
		globalLogger.Error(err)
		return err
	}

	err = container.Singleton(func() models.IAuthService {
		return services.NewAuthService()
	})
	if err != nil {
		globalLogger.Error(err)
		return err
	}

	err = container.Singleton(func() models.IPostService {
		return services.NewPostService()
	})

	if err != nil {
		globalLogger.Error(err)
		return err
	}

	return nil
}
