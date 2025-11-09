package config

import (
	"log"

	"github.com/inidaname/arkade-vault/api-gateway/pkg/types"
	"github.com/joho/godotenv"
)

// Global Config

// LoadConfig loads environment variables and sets the configuration
func LoadConfig() types.Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Unable to load .env file")
	}

	appConfig := types.Config{}

	return appConfig
}