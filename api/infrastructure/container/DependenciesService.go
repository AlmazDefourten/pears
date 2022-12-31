package container

import (
	"github.com/AlmazDefourten/goapp/interface/handler"
	models "github.com/AlmazDefourten/goapp/models"
	log "github.com/sirupsen/logrus"
	viper "github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewContainer Constructor for Container
func NewContainer(db *gorm.DB, configurator *viper.Viper) models.Container {
	return models.Container{
		AppConnection:  db,
		ConfigProvider: configurator,
	}
}

func NewServiceContainer(userService models.IUserService) models.ServiceContainer {
	return models.ServiceContainer{
		UserService: userService,
	}
}

// NewConnection Constructor that initialize new connection
func NewConnection(viperInit models.Configurator) *gorm.DB {
	// Get a connstring from config
	connstring := viperInit.GetString("connection")

	// Open new connection
	db, err := gorm.Open(postgres.Open(connstring), &gorm.Config{})

	if err != nil {
		log.WithFields(log.Fields{
			"error":      err.Error(),
			"connstring": db.Config,
		}).Fatal("Failed to connect to database")
	}

	return db
}

type HandlerContainer struct {
	UserInfoHandler *handler.UserInfoHandler
}

func NewHandlerContainer(userHandler *handler.UserInfoHandler) HandlerContainer {
	return HandlerContainer{
		UserInfoHandler: userHandler,
	}
}
