package models

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// Container Struct for store global variables for app
type Container struct {
	AppConnection  *gorm.DB
	ConfigProvider *viper.Viper
}

// ServiceContainer for store services singleton
type ServiceContainer struct {
	UserService IUserService
}
