package database

import (
	"main/database/model"

	"gorm.io/gorm"
)

var modelsList = []interface{}{
	&model.User{},
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(modelsList...)
}
