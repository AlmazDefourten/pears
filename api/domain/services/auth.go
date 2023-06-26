package services

import (
	"github.com/AlmazDefourten/goapp/infra/logger_instance"
	"github.com/AlmazDefourten/goapp/models/requests_models"
	"github.com/AlmazDefourten/goapp/models/user_models"
	"github.com/AlmazDefourten/goapp/models/util_adapters"
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

const (
	defaultUserExistFlagValue = false
	regAnswerIfUserExist      = "Пользователь с таким именем уже существует"
	regAnswerSuccessful       = "Регистрация прошла успешно"
	authAnswerSuccessful      = "Авторизация прошла успешно"
	authAnswerFailed          = "Неправильный логин или пароль"
)

func (authService *AuthService) CheckIfUserExist(login string) (bool, error) {
	var res []user_models.User
	var db gorm.DB
	err := container.Resolve(&db)
	if err != nil {
		logger_instance.ServiceLogger.Error(err)
		return defaultUserExistFlagValue, err
	}
	//request := db.Model(&models.User{}).First(&res, "login = ?", login)
	request := db.Model(&user_models.User{}).Where("login = ?", login).Find(&res)

	if request.Error != nil && !errors.Is(request.Error, gorm.ErrRecordNotFound) {
		logger_instance.ServiceLogger.Error(request.Error)
		return defaultUserExistFlagValue, request.Error
	}
	if request.RowsAffected > 0 {
		return true, nil
	}
	return false, nil
}

func (authService *AuthService) Registration(user *user_models.User) (bool, string) {
	var c util_adapters.Configurator
	err := container.Resolve(&c)
	if err != nil {
		logger_instance.ServiceLogger.Error(err)
		return false, requests_models.StandardAnswerOnError
	}
	var db gorm.DB
	err = container.Resolve(&db)
	if err != nil {
		logger_instance.ServiceLogger.Error(err)
		return false, requests_models.StandardAnswerOnError
	}
	isUserExists, err := authService.CheckIfUserExist(user.Login)
	if err != nil {
		logger_instance.ServiceLogger.Error(err)
		return false, requests_models.StandardAnswerOnError
	}
	if isUserExists {
		return false, regAnswerIfUserExist
	} else {
		user.Password = hashPassword(user.Password,
			c.GetString("passwordSalt"),
			c.GetInt("hashingCost"))
		request := db.Create(&user)
		if request.Error != nil {
			logger_instance.ServiceLogger.Error(err)
			return false, requests_models.StandardAnswerOnError
		}
	}
	return true, regAnswerSuccessful
}

func (authService *AuthService) Authorization(login string, password string) (bool, *user_models.Tokens, string) {
	var db gorm.DB
	err := container.Resolve(&db)
	if err != nil {
		logger_instance.ServiceLogger.Error(err)
		return false, nil, requests_models.StandardAnswerOnError
	}

	var c util_adapters.Configurator
	err = container.Resolve(&c)
	if err != nil {
		logger_instance.ServiceLogger.Error(err)
		return false, nil, requests_models.StandardAnswerOnError
	}

	var jwtService user_models.IJWTService
	err = container.Resolve(&jwtService)
	if err != nil {
		logger_instance.ServiceLogger.Error(err)
		return false, nil, requests_models.StandardAnswerOnError
	}

	var user user_models.User
	err = db.Where("login = ?", login).Find(&user).Error
	//err = db.First(&user, "login = ?", login).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil, requests_models.StandardAnswerOnError
		}
	}

	if checkPasswordHash(password, user.Password, c.GetString("passwordSalt")) {
		jwtToken, err := jwtService.SignIn(user.Id)
		if err != nil {
			logger_instance.ServiceLogger.Error(err)
			return false, nil, authAnswerFailed
		}
		return true, jwtToken, authAnswerSuccessful
	} else {
		return false, nil, authAnswerFailed
	}
}

func (authService *AuthService) AuthCheck(token string) (bool, string) {
	var c util_adapters.Configurator
	err := container.Resolve(&c)

	username, err := ParseToken(token, []byte(c.GetString("jwt.signing_key")))
	if err != nil {
		logger_instance.ServiceLogger.Error(err)
		return false, ""
	}
	return true, username
}

func (authService *AuthService) RefreshCheck(token string) (bool, *user_models.Tokens) {
	var db gorm.DB
	err := container.Resolve(&db)
	if err != nil {
		panic(err)
	}

	var user user_models.User
	err = db.First(&user, "refresh_token = ?", token).Error
	if err != nil {
		return false, nil
	}

	var jwtService user_models.IJWTService
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
		logger_instance.ServiceLogger.Error(err)
	}
	return string(bytes)
}

func checkPasswordHash(password, hash string, passwordSalt string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password+passwordSalt))
	return err == nil
}
