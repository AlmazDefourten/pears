package container

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	viper "github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Struct for store global variables for app
type Container struct {
	AppConnection  *gorm.DB
	ConfigProvider *viper.Viper
}

// Constructor for Container
func NewContainer(db *gorm.DB, configurator *viper.Viper) Container {
	return Container{
		AppConnection:  db,
		ConfigProvider: configurator,
	}
}

// Constructor that initialize new viper configurator
func NewViperConfigurator() *viper.Viper {
	viper_init := viper.New()

	viper_init.SetConfigName("appconfig") // name of config file (without extension)
	viper_init.SetConfigType("json")      // REQUIRED if the config file does not have the extension in the name
	viper_init.AddConfigPath(".")         // optionally look for config in the working directory

	err := viper_init.ReadInConfig() // Find and read the config file
	if err != nil {                  // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	return viper_init
}

// Constructor that initialize new connection
func NewConnection(viper_init *viper.Viper) *gorm.DB {
	// Get a connstring from config
	connstring := viper_init.GetString("connection")

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
