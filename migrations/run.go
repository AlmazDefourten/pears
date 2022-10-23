package migrations

import (
	model "github.com/AlmazDefourten/goapp/models"
	"gorm.io/gorm"
)

// Function for migrate all ORM entities only
func RunBaseMigration(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
}
