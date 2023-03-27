package services

import (
	"github.com/AlmazDefourten/goapp/infrastructure/loggerInstance"
	"github.com/AlmazDefourten/goapp/models"
	"github.com/golobby/container/v3"
	"github.com/kataras/iris/v12/x/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//TODO: возвращать error в методах
type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (authService *AuthService) CheckIfUserExist(login string) bool {
	var res []models.User
	var db gorm.DB
	err := container.Resolve(&db)
	if err != nil {
		loggerInstance.ServiceLogger.Error(err)
		panic(err)
	}
	request := db.Model(&models.User{}).First(&res, "login = ?", login)
	if request.Error != nil {
		loggerInstance.ServiceLogger.Error(request.Error)
	}
	if len(res) > 0 {
		return true
	}
	return false
}

func (authService *AuthService) Registration(user *models.User) bool {
	var c models.Configurator
	err := container.Resolve(&c)
	if err != nil {
		loggerInstance.ServiceLogger.Error(err)
		panic(err)
	}
	var db gorm.DB
	err = container.Resolve(&db)
	if err != nil {
		loggerInstance.ServiceLogger.Error(err)
		panic(err)
	}
	if authService.CheckIfUserExist(user.Login) {
		// some info in callback idk that user already exists
		return false
	} else {
		user.Password = hashPassword(user.Password,
			c.GetString("passwordSalt"),
			c.GetInt("hashingCost"))
		request := db.Create(&user)
		if request.Error != nil {
			// log error lol
			return false
		}
	}
	return true
}

func (authService *AuthService) Authorization(login string, password string) (bool, *models.Tokens) {
	var db gorm.DB
	err := container.Resolve(&db)
	if err != nil {
		loggerInstance.ServiceLogger.Error(err)
		panic(err)
	}

	var c models.Configurator
	err = container.Resolve(&c)
	if err != nil {
		loggerInstance.ServiceLogger.Error(err)
		panic(err)
	}

	var jwtService models.IJWTService
	err = container.Resolve(&jwtService)
	if err != nil {
		loggerInstance.ServiceLogger.Error(err)
		panic(err)
	}

	var user models.User
	err = db.First(&user, "login = ?", login).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
	}
	if checkPasswordHash(password, user.Password, c.GetString("passwordSalt")) {
		jwtToken, err := jwtService.SignIn(login)
		if err != nil {
			loggerInstance.ServiceLogger.Error(err)
		}
		return true, jwtToken
	} else {
		return false, nil
	}
}

func (authService *AuthService) AuthCheck(token string) (bool, string) {
	var c models.Configurator
	err := container.Resolve(&c)

	username, err := ParseToken(token, []byte(c.GetString("jwt.signing_key")))
	if err != nil {
		loggerInstance.ServiceLogger.Error(err)
		return false, ""
	}
	return true, username
}

func (authService *AuthService) RefreshCheck(token string) (bool, *models.Tokens) {
	var db gorm.DB
	err := container.Resolve(&db)
	if err != nil {
		panic(err)
	}

	var user models.User
	err = db.First(&user, "refresh_token = ?", token).Error
	if err != nil {
		return false, nil
	}

	var jwtService models.IJWTService
	err = container.Resolve(&jwtService)
	if err != nil {
		panic(err)
	}

	tokens, err := jwtService.ValidateAndRefreshTokens(token)

	if err != nil {
		return false, nil
	}

	return true, tokens
}

func hashPassword(password string, passwordSalt string, hashingCost int) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password+passwordSalt), hashingCost)
	if err != nil {
		loggerInstance.ServiceLogger.Error(err)
	}
	return string(bytes)
}

func checkPasswordHash(password, hash string, passwordSalt string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password+passwordSalt))
	return err == nil
}
