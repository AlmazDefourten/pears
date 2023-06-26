package user_repository

import (
	models "github.com/AlmazDefourten/goapp/models/user_models"
	"github.com/golobby/container/v3"
	"gorm.io/gorm"
)

type UserRepo struct {
}

func NewUserRepository() *UserRepo {
	return &UserRepo{}
}

func (repo UserRepo) List() []models.User {
	var db gorm.DB
	err := container.Resolve(&db)
	if err != nil {
		return nil
	}
	var users []models.User
	db.Find(&users)
	return users
}

func (repo UserRepo) Get(id int) (models.User, bool, string) {
	//TODO: реализовать методы
	return models.User{}, false, ""
}

func (repo UserRepo) Delete(id int) (bool, string) {
	return false, ""
}

func (repo UserRepo) Update(user models.User) (bool, string) {
	return false, ""
}

func (repo UserRepo) Create(models.User) (bool, string) {
	return false, ""
}
