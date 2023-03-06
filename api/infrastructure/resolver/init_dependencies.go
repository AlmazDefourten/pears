package resolver

import (
	"github.com/AlmazDefourten/goapp/infrastructure/configurator"
	"github.com/AlmazDefourten/goapp/infrastructure/data_adapter/connection"
	"github.com/AlmazDefourten/goapp/models"
	"github.com/AlmazDefourten/goapp/services"
	"github.com/golobby/container/v3"
	"gorm.io/gorm"
)

// InitializeContainer with global app dependencies -
// Connection, configurator, etc...
func InitializeContainer() error {
	err := container.Singleton(func() models.Configurator {
		return configurator.NewViperConfigurator()
	})
	if err != nil {
		//logging here
		return err
	}
	var c models.Configurator
	err = container.Resolve(&c)
	if err != nil {
		//logging here
		return err
	}

	err = container.Singleton(func() gorm.DB {
		return *connection.NewGormConnection(c)
	})
	if err != nil {
		//logging here
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
		//logging here
		return err
	}

	err = container.Singleton(func() models.IJWTService {
		return services.NewJWTService()
	})
	if err != nil {
		//logging here
		return err
	}

	err = container.Singleton(func() models.IAuthService {
		return services.NewAuthService()
	})
	if err != nil {
		//logging here
		return err
	}

	err = container.Singleton(func() models.IPostService {
		return services.NewPostService()
	})
	if err != nil {
		//logging here
		return err
	}
	return nil
}
