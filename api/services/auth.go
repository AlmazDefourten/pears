package services

import (
	"fmt"
	"github.com/AlmazDefourten/goapp/models"
	"github.com/AlmazDefourten/goapp/models/container_models"
	"github.com/kataras/iris/v12/x/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	Container  *container_models.Container
	JWTService models.IJWTService
}

func NewAuthService(container *container_models.Container, jwtService models.IJWTService) *AuthService {
	return &AuthService{
		Container:  container,
		JWTService: jwtService,
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
		user.Password = hashPassword(user.Password,
			authService.Container.ConfigProvider.GetString("passwordSalt"),
			authService.Container.ConfigProvider.GetInt("hashingCost"))
		request := authService.Container.AppConnection.Create(&user)
		if request.Error != nil {
			// log error lol
			return false
		}
	}
	return true
}

func (authService *AuthService) Authorization(login string, password string) (bool, *models.Tokens) {
	var user models.User
	err := authService.Container.AppConnection.First(&user, "login = ?", login).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
	}
	if checkPasswordHash(password, user.Password, authService.Container.ConfigProvider.GetString("passwordSalt")) {
		jwtToken, err := authService.JWTService.SignIn(login)
		if err != nil {
			fmt.Println(err)
			// logging here lol
		}
		return true, jwtToken
	} else {
		return false, nil
	}
}

func (authService *AuthService) AuthCheck(token string) (bool, string) {
	username, err := ParseToken(token, []byte(authService.Container.ConfigProvider.GetString("jwt.signing_key")))
	if err != nil {
		fmt.Println(err)
		// logging here lol
		return false, ""
	}
	return true, username
}

func hashPassword(password string, passwordSalt string, hashingCost int) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password+passwordSalt), hashingCost)
	if err != nil {
		println(err)
		// logging here lol
	}
	return string(bytes)
}

func checkPasswordHash(password, hash string, passwordSalt string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password+passwordSalt))
	return err == nil
}
