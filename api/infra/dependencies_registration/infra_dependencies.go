package dependencies_registration

import (
	"github.com/AlmazDefourten/goapp/infra/configurator"
	"github.com/AlmazDefourten/goapp/infra/data_adapter/connection"
	"github.com/AlmazDefourten/goapp/infra/logger_instance"
	"github.com/AlmazDefourten/goapp/models/util_adapters"
	"github.com/AlmazDefourten/goapp/pkg/logging/loggers"
	"github.com/AlmazDefourten/goapp/pkg/logging/resolvers"
	"github.com/golobby/container/v3"
	"gorm.io/gorm"
)

// InitializeContainer with global app dependencies -
// Connection, configurator, etc...
func InitializeContainer() error {
	err := container.NamedSingleton("GlobalLogger", func() util_adapters.Logger {
		return loggers.Init(resolvers.GlobalLogger)
	})
	err = container.NamedResolve(&logger_instance.GlobalLogger, "GlobalLogger")
	if err != nil {
		return err
	}
	err = container.NamedSingleton("ServiceLogger", func() util_adapters.Logger {
		return loggers.Init(resolvers.ServiceLogger)
	})
	err = container.NamedResolve(&logger_instance.ServiceLogger, "ServiceLogger")
	if err != nil {
		logger_instance.GlobalLogger.Error(err)
		return err
	}
	if err != nil {
		logger_instance.GlobalLogger.Error(err)
		return err
	}

	err = container.Singleton(func() util_adapters.Configurator {
		return configurator.NewViperConfigurator()
	})

	if err != nil {
		logger_instance.GlobalLogger.Error(err)
		return err
	}
	var c util_adapters.Configurator
	err = container.Resolve(&c)
	if err != nil {
		logger_instance.GlobalLogger.Error(err)
		return err
	}

	err = container.Singleton(func() gorm.DB {
		return *connection.NewGormConnection(c)
	})
	if err != nil {
		logger_instance.GlobalLogger.Error(err)
		return err
	}
	return nil
}
