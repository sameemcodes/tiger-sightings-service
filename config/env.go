package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("unable to load .env file")
	}
}

func GetEnv(key string, defaultVal string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Printf("Environment variable %s not set, using default value: %s", key, defaultVal)
		return defaultVal
	}
	return value
}
