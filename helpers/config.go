package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
    PORT                 string
    SWAGGER_ENABLED      string
    BASIC_AUTH_USER      string
    BASIC_AUTH_PASSWORD  string
    JWT_SECRET           string
}

// Define a global variable to hold the config
var AppConfig Config

func LoadConfig() {
    // Load .env file
    err := godotenv.Load()
    if err != nil {
        fmt.Println("[ERROR] Error loading .env file:", err)
    }

    // Set values to AppConfig
    AppConfig = Config{
        PORT:                 os.Getenv("PORT"),
        SWAGGER_ENABLED:      os.Getenv("SWAGGER_ENABLED"),
        BASIC_AUTH_USER:      os.Getenv("BASIC_AUTH_USER"),
        BASIC_AUTH_PASSWORD:  os.Getenv("BASIC_AUTH_PASSWORD"),
        JWT_SECRET:           os.Getenv("JWT_SECRET"),
    }
}