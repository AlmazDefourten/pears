package container_models

import (
	"github.com/AlmazDefourten/goapp/interface/handler"
	"github.com/AlmazDefourten/goapp/models"
	"gorm.io/gorm"
)

type HandlerContainer struct {
	UserInfoHandler *handler.UserInfoHandler
	AuthHandler     *handler.AuthHandler
	PostHandler     *handler.PostHandler
}

// Container Struct for store global variables for app
type Container struct {
	AppConnection  *gorm.DB
	ConfigProvider models.Configurator
}

// ServiceContainer for store services singleton
type ServiceContainer struct {
	UserService models.IUserService
	JWTService  models.IJWTService
	AuthService models.IAuthService
	PostService models.IPostService
}
