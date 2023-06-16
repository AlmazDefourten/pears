package repo_models

import models "github.com/AlmazDefourten/goapp/models/user_models"

type UserRepository interface {
	List() []models.User
	Create(models.User) (bool, string)
	Update(models.User) (bool, string)
	Get(id int) (models.User, bool, string)
	Delete(id int) (bool, string)
}
