package configuration

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
}

func GetConfigValue(key string) string {
	return os.Getenv(key)
}
