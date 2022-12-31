package models

import (
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

// IUserService interface for operations with Users
type IUserService interface {
	List() []User
}
