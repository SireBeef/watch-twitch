package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ClientID        string
	UserID          string
	UserAccessToken string
	BrowserToken    string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		ClientID:        os.Getenv("CLIENT_ID"),
		UserID:          os.Getenv("USER_ID"),
		UserAccessToken: os.Getenv("USER_ACCESS_TOKEN"),
		BrowserToken:    os.Getenv("BROWSER_AUTH_TOKEN"),
	}
}
