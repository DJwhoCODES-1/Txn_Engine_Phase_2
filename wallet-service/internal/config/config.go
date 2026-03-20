package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GRPCPort string
	MongoURI string
	MongoDB  string
}

func Load() *Config {
	_ = godotenv.Load()

	cfg := &Config{
		GRPCPort: getEnv("GRPC_PORT", "50052"),
		MongoURI: getEnv("MONGO_URI", "mongodb://localhost:27017"),
		MongoDB:  getEnv("MONGO_DB", "wallet-db"),
	}

	log.Println("Wallet Service Config Loaded")
	return cfg
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
