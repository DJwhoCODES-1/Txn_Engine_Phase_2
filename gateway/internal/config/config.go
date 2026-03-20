package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port            string
	AdminServiceURL string
}

func Load() *Config {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	cfg := &Config{
		Port:            getEnv("PORT", "8080"),
		AdminServiceURL: getEnv("ADMIN_SERVICE_URL", "localhost:50051"),
	}

	log.Println("Config loaded")
	return cfg
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
