package database

import (
	"main/database/models"

	"gorm.io/gorm"
)

var modelsList = []interface{}{
	&models.User{},
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(modelsList...)
}
