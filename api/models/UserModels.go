package models

import (
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

// User Struct for User entity
type User struct {
	gorm.Model
	Id       int
	Login    string
	Password string
	Name     string
	Age      int
}

// Claims Struct for JWT Claims
type Claims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

// IJWTService interface for operations with JWT
type IJWTService interface {
	SignIn(username, password string) (string, error)
	CheckToken(accessToken string, signingKey []byte) (string, error)
}

// IUserService interface for operations with Users
type IUserService interface {
	List() []User
}
