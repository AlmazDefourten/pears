package models

import (
	gorm "gorm.io/gorm"
)

// Struct for User entity
type User struct {
	gorm.Model
	Name string
	Age  int
}
