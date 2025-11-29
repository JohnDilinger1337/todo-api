package database

import (
	"fmt"
	"log"
	"main/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(cfg *config.Config) error {
	var err error

	DB, err = gorm.Open(sqlite.Open(cfg.DBPath), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to SQLite:", err)
		return fmt.Errorf("Failed to connect to SQLite: %s", err)
	}

	log.Println("SQLite connected:", cfg.DBPath)
	return nil
}
