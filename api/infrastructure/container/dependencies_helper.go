package container

import (
	"github.com/AlmazDefourten/goapp/interface/handler"
	models "github.com/AlmazDefourten/goapp/models"
	"github.com/AlmazDefourten/goapp/models/container_models"
	viper "github.com/spf13/viper"
	"gorm.io/gorm"
)

// NewContainer Constructor for Container
func NewContainer(db *gorm.DB, configurator *viper.Viper) container_models.Container {
	return container_models.Container{
		AppConnection:  db,
		ConfigProvider: configurator,
	}
}

func NewServiceContainer(userService models.IUserService, jwtService models.IJWTService) container_models.ServiceContainer {
	return container_models.ServiceContainer{
		UserService: userService,
		JWTService:  jwtService,
	}
}

func NewHandlerContainer(userHandler *handler.UserInfoHandler) container_models.HandlerContainer {
	return container_models.HandlerContainer{
		UserInfoHandler: userHandler,
	}
}
