package utils

import (
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	Port   string 
	DBConnString string
}

// NewConfig creates a new instance of Config.

// If in development, loads env vars from .env file.
func NewConfig() *config {
	if os.Getenv("ENV") != "production" {
		_ = godotenv.Load(".env")
	}
 
	return &config{ 
		Port:   getEnvVar("PORT", "3000 "),
		DBConnString: getEnvVar("DBConnString", "goDB"),
	}
}

// getEnvVar searches for a given key or return the fallback of key doesn't exist.
func getEnvVar(key, fallback string) string {
	if val, exist := os.LookupEnv(key); exist {
		return val
	}
	return fallback
} 