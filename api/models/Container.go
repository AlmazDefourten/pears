package models

import (
	"gorm.io/gorm"
)

// Container Struct for store global variables for app
type Container struct {
	AppConnection  *gorm.DB
	ConfigProvider Configurator
}

// ServiceContainer for store services singleton
type ServiceContainer struct {
	UserService IUserService
}
