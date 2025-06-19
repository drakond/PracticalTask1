package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	AppName    string
	LogLevel   string
	ServerName string
}

func Load() *Config {
	// Загружаем .env если есть
	_ = godotenv.Load(".env")

	return &Config{
		Port:       getEnv("PORT", ":3000"),
		AppName:    getEnv("APP_NAME", "prtask1"),
		LogLevel:   getEnv("LOG_LEVEL", "info"),
		ServerName: getEnv("SERVER_NAME", "prtask1"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
