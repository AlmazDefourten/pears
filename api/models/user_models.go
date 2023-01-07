package models

import (
	"gorm.io/gorm"
)

// User Struct for User entity
type User struct {
	gorm.Model
	Name string
	Age  int
}

// IUserService interface for operations with Users
type IUserService interface {
	List() []User
}
