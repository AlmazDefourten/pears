package connection

import (
	"github.com/AlmazDefourten/goapp/infrastructure/loggerInstance"
	"github.com/AlmazDefourten/goapp/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewConnection Constructor that initialize new connection
func NewGormConnection(viperInit models.Configurator) *gorm.DB {
	// Get a connstring from config
	connstring := viperInit.GetString("connection")

	// Open new connection
	db, err := gorm.Open(postgres.Open(connstring), &gorm.Config{})

	if err != nil {
		loggerinstance.GlobalLogger.Error(err)
	}

	return db
}
