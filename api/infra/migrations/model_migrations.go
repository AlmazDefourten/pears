package migrations

import (
	"github.com/AlmazDefourten/goapp/infra/logger_instance"
	"github.com/AlmazDefourten/goapp/models/post_models"
	"github.com/AlmazDefourten/goapp/models/tag_models"
	model "github.com/AlmazDefourten/goapp/models/user_models"
	"github.com/golobby/container/v3"
	"gorm.io/gorm"
)

// RunBaseMigration for migrate all ORM entities only
func RunBaseMigration() {
	var db gorm.DB
	err := container.Resolve(&db)
	if err != nil {
		logger_instance.GlobalLogger.Error(err)
		panic(err)
	}
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		logger_instance.GlobalLogger.Error(err)
		panic(err)
	}
	err = db.AutoMigrate(&post_models.Post{})
	if err != nil {
		logger_instance.GlobalLogger.Error(err)
		panic(err)
	}
	err = db.AutoMigrate(&tag_models.Tag{})
	if err != nil {
		logger_instance.GlobalLogger.Error(err)
		panic(err)
	}
	err = db.AutoMigrate(&post_models.PostTags{})
	if err != nil {
		logger_instance.GlobalLogger.Error(err)
		panic(err)
	}
}
