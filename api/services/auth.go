package services

import (
	"github.com/AlmazDefourten/goapp/models"
	"github.com/AlmazDefourten/goapp/models/container_models"
	"github.com/kataras/iris/v12/x/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	passwordSalt = "ASdlamsdpqwekAvnaQRIBFQHYWI1523hjsdhASHDSUDQWEklasdaousdgBVIVQOPWE"
	hashingCost  = 14
)

type AuthService struct {
	Container *container_models.Container
}

func NewAuthService(container *container_models.Container) *AuthService {
	return &AuthService{
		Container: container,
	}
}

func (authService *AuthService) CheckIfUserExist(login string) bool {
	var res []models.User
	request := authService.Container.AppConnection.Model(&models.User{}).First(&res, "login = ?", login)
	if request.Error != nil {
		// logging and debug
	}
	if len(res) > 0 {
		return true
	}
	return false
}

func (authService *AuthService) Registration(user *models.User) bool {
	if authService.CheckIfUserExist(user.Login) {
		// some info in callback idk that user already exists
		return false
	} else {
		user.Password = hashPassword(user.Password)
		request := authService.Container.AppConnection.Create(&user)
		if request.Error != nil {
			// log error lol
			return false
		}
	}
	return true
}

func (authService *AuthService) Authorization(login string, password string) bool {
	var user models.User
	err := authService.Container.AppConnection.First(&user, "login = ?", login).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false
		}
	}
	if checkPasswordHash(password, user.Password) {
		return true
	} else {
		return false
	}
}

func hashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password+passwordSalt), hashingCost)
	if err != nil {
		println(err)
		// logging here lol
	}
	return string(bytes)
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password+passwordSalt))
	return err == nil
}
