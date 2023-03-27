package models

import (
	"github.com/dgrijalva/jwt-go"
)

// User Struct for User entity
//TODO: Write examople refresh token
type User struct {
	Id            int    `json:"id"`
	Login         string `json:"login" example:"ivan.petrov@mail.ru"`
	Nick          string `json:"nick" example:"PedanticCow"`
	Password      string `json:"password" example:"mypaSSword123876"`
	Name          string `json:"name" example:"Andrey"`
	Age           int    `json:"age" example:"24"`
	Refresh_token string `json:"refresh_token" example:""`
}

// Claims Struct for JWT Claims
type Claims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

// IJWTService interface for operations with JWT
type IJWTService interface {
	SignIn(username string) (*Tokens, error)
	ValidateAndRefreshTokens(refresh_token string) (*Tokens, error)
}

// IUserService interface for operations with Users
type IUserService interface {
	List() []User
}

// tokens struct for JWT
type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
