package models

import (
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

// User Struct for User entity
type User struct {
	gorm.Model
	Id       int    `json:"id"`
	Login    string `json:"login"`
	Nick     string `json:"nick"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
}

// Claims Struct for JWT Claims
type Claims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

// IJWTService interface for operations with JWT
type IJWTService interface {
	SignIn(username, password string) (string, error)
}

// IUserService interface for operations with Users
type IUserService interface {
	List() []User
}
