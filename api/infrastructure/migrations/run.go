package migrations

import (
	"github.com/AlmazDefourten/goapp/infrastructure/loggerInstance"
	model "github.com/AlmazDefourten/goapp/models"
	"github.com/golobby/container/v3"
	"gorm.io/gorm"
)

// RunBaseMigration for migrate all ORM entities only
func RunBaseMigration() {
	var db gorm.DB
	err := container.Resolve(&db)
	if err != nil {
		loggerInstance.GlobalLogger.Error(err)
		panic(err)
	}
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		loggerInstance.GlobalLogger.Error(err)
		panic(err)
	}
	err = db.AutoMigrate(&model.Post{})
	if err != nil {
		loggerInstance.GlobalLogger.Error(err)
		panic(err)
	}
	err = db.AutoMigrate(&model.Tag{})
	if err != nil {
		loggerInstance.GlobalLogger.Error(err)
		panic(err)
	}
	err = db.AutoMigrate(&model.PostTags{})
	if err != nil {
		loggerInstance.GlobalLogger.Error(err)
		panic(err)
	}
}
