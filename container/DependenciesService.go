package container

import (
	"fmt"
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

// NewViperConfigurator Constructor that initialize new viper configurator
func NewViperConfigurator() *viper.Viper {
	viperInit := viper.New()

	viperInit.SetConfigName("appconfig") // name of config file (without extension)
	viperInit.SetConfigType("json")      // REQUIRED if the config file does not have the extension in the name
	viperInit.AddConfigPath(".")         // optionally look for config in the working directory

	err := viperInit.ReadInConfig() // Find and read the config file
	if err != nil {                 // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	return viperInit
}

// NewConnection Constructor that initialize new connection
func NewConnection(viperInit *viper.Viper) *gorm.DB {
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
