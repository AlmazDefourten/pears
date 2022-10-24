package models

import (
	"github.com/spf13/viper"
	gorm "gorm.io/gorm"
)

// Struct for store global variables for app
type Container struct {
	AppConnection  *gorm.DB
	ConfigProvider *viper.Viper
}

type Services struct {
	UserService IUserService
}
