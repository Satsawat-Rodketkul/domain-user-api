package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	rootPath, err := filepath.Abs("../.env")
	if err != nil {
		log.Fatal("Error getting .env file path")
	}

	err = godotenv.Load(rootPath)
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
}

func GetConfigValue(key string) string {
	return os.Getenv(key)
}
