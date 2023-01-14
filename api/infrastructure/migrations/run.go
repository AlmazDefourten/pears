package migrations

import (
	model "github.com/AlmazDefourten/goapp/models"
	"gorm.io/gorm"
)

// RunBaseMigration for migrate all ORM entities only
func RunBaseMigration(db *gorm.DB) {
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		// there is logging
		panic(err)
	}
}
