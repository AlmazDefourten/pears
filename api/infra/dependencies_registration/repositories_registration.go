package dependencies_registration

import (
	"github.com/AlmazDefourten/goapp/infra/data_adapter/repository/user_repository"
	"github.com/AlmazDefourten/goapp/infra/logger_instance"
	"github.com/AlmazDefourten/goapp/models/repo_models"
	"github.com/golobby/container/v3"
)

func InitializeRepositories() error {
	err := container.Singleton(func() repo_models.UserRepository {
		return user_repository.NewUserRepository()
	})
	if err != nil {
		logger_instance.GlobalLogger.Error("Ошибка при регистрации репоизтория UserRepository")
		return err
	}
	return nil
}
