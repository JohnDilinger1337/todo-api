package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port    string
	GinMode string
	DBPath  string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = "debug"
	}
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "data.sqlite"
	}

	return &Config{Port: port, GinMode: ginMode, DBPath: dbPath}

}
