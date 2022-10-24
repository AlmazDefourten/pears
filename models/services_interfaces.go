package models

// Service for operations with Users
type IUserService interface {
	List() []User
}
