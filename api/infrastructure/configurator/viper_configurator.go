package configurator

import (
	"fmt"
	"github.com/spf13/viper"
)

// NewViperConfigurator Constructor that initialize new viper configurator
func NewViperConfigurator() *viper.Viper {
	viperInit := viper.New()

	viperInit.SetConfigName(ConfigName) // name of config file (without extension)
	viperInit.SetConfigType(ConfigType) // REQUIRED if the config file does not have the extension in the name
	viperInit.AddConfigPath(ConfigPath) // optionally look for config in the working directory

	err := viperInit.ReadInConfig() // Find and read the config file
	if err != nil {                 // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	return viperInit
}
