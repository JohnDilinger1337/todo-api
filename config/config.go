package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName      string
	Port         string
	GinMode      string
	DBPath       string
	JWTSecret    string
	JWTExpiresAt string
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
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "supersecretkey"
	}
	jwtExpiresAtStr := os.Getenv("JWT_EXPIRES_AT")
	if jwtExpiresAtStr == "" {
		jwtExpiresAtStr = "72h"
	}
	appName := os.Getenv("APP_NAME")
	if appName == "" {
		appName = "Todo API"
	}

	return &Config{Port: port, GinMode: ginMode, DBPath: dbPath, JWTSecret: jwtSecret, JWTExpiresAt: jwtExpiresAtStr, AppName: appName}

}
