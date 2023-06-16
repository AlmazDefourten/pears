package connection

import (
	"github.com/AlmazDefourten/goapp/infra/logger_instance"
	"github.com/AlmazDefourten/goapp/models/util_adapters"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewGormConnection Constructor that initialize new connection
func NewGormConnection(viperInit util_adapters.Configurator) *gorm.DB {
	// Get a connstring from config
	connstring := viperInit.GetString("connection")

	// Open new connection
	db, err := gorm.Open(postgres.Open(connstring), &gorm.Config{})

	if err != nil {
		logger_instance.GlobalLogger.Error(err)
	}

	return db
}
