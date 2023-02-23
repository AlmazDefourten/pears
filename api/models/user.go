package models

import (
	"github.com/dgrijalva/jwt-go"
)

// User Struct for User entity
type User struct {
	Id       int    `json:"id"`
	Login    string `json:"login" example:"ivan.petrov@mail.ru"`
	Nick     string `json:"nick" example:"PedanticCow"`
	Password string `json:"password" example:"mypaSSword123876"`
	Name     string `json:"name" example:"Andrey"`
	Age      int    `json:"age" example:"24"`
}

// Claims Struct for JWT Claims
type Claims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

// IJWTService interface for operations with JWT
type IJWTService interface {
	SignIn(username string) (string, error)
}

// IUserService interface for operations with Users
type IUserService interface {
	List() []User
}
