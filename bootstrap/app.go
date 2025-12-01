package bootstrap

import (
	"main/config"
	"main/database"

	"gorm.io/gorm"
)

type App struct {
	DB *gorm.DB
}

func InitializeApp(cfg *config.Config) (*gorm.DB, error) {
	db, err := database.Connect(cfg)
	if err != nil {
		return nil, err
	}

	if err := database.Migrate(db); err != nil {
		return nil, err
	}

	return db, nil
}
