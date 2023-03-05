package migrations

import (
	model "github.com/AlmazDefourten/goapp/models"
	"github.com/golobby/container/v3"
	"gorm.io/gorm"
)

// RunBaseMigration for migrate all ORM entities only
func RunBaseMigration() {
	var db gorm.DB
	err := container.Resolve(&db)
	if err != nil {
		//logging here
		panic(err)
	}
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		// there is logging
		panic(err)
	}
	err = db.AutoMigrate(&model.Post{})
	if err != nil {
		// there is logging
		panic(err)
	}
	err = db.AutoMigrate(&model.Tag{})
	if err != nil {
		// there is logging
		panic(err)
	}
	err = db.AutoMigrate(&model.PostTags{})
	if err != nil {
		// there is logging
		panic(err)
	}
}
