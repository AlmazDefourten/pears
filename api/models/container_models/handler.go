package container_models

import (
	"github.com/AlmazDefourten/goapp/interface/handler"
	"github.com/AlmazDefourten/goapp/models/auth_models"
	"github.com/AlmazDefourten/goapp/models/post_models"
	"github.com/AlmazDefourten/goapp/models/user_models"
	"github.com/AlmazDefourten/goapp/models/util_adapters"
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
	ConfigProvider util_adapters.Configurator
}

// ServiceContainer for store services singleton
type ServiceContainer struct {
	UserService user_models.IUserService
	JWTService  user_models.IJWTService
	AuthService auth_models.IAuthService
	PostService post_models.IPostService
}
