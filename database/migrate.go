package database

import (
	"main/database/model"

	"gorm.io/gorm"
)

var modelsList = []interface{}{
	&model.User{},
	&model.TodoItem{},
	&model.TodoCategory{},
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(modelsList...)
}
