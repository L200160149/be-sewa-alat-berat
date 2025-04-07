package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func GetEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
