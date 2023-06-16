package dependencies_registration

import (
	"github.com/AlmazDefourten/goapp/domain/services"
	"github.com/AlmazDefourten/goapp/infra/logger_instance"
	"github.com/AlmazDefourten/goapp/models"
	"github.com/golobby/container/v3"
)

// RegisterServices - decomposition ServiceContainer to services
func RegisterServices() error {
	err := container.Singleton(func() models.IUserService {
		return services.NewUserService()
	})
	if err != nil {
		logger_instance.GlobalLogger.Error(err)
		return err
	}

	err = container.Singleton(func() models.IJWTService {
		return services.NewJWTService()
	})
	if err != nil {
		logger_instance.GlobalLogger.Error(err)
		return err
	}

	err = container.Singleton(func() models.IAuthService {
		return services.NewAuthService()
	})
	if err != nil {
		logger_instance.GlobalLogger.Error(err)
		return err
	}

	err = container.Singleton(func() models.IPostService {
		return services.NewPostService()
	})

	if err != nil {
		logger_instance.GlobalLogger.Error(err)
		return err
	}

	return nil
}
