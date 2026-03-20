package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GRPCPort  string
	JWTSecret string
	MongoURI  string
}

func Load() *Config {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	cfg := &Config{
		GRPCPort:  getEnv("GRPC_PORT", "50051"),
		JWTSecret: getEnv("JWT_SECRET", "super-secret"),
		MongoURI:  getEnv("MONGO_URI", "mongodb://localhost:27017"),
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
