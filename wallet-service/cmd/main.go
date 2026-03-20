package main

import (
	"log"

	"txn-engine-phase-2/wallet-service/internal/config"
	"txn-engine-phase-2/wallet-service/internal/database"
	"txn-engine-phase-2/wallet-service/internal/server"
)

func main() {
	cfg := config.Load()

	db := database.ConnectMongo(cfg.MongoURI, cfg.MongoDB)

	s := server.NewGRPCServer(cfg, db)

	log.Printf("Wallet service running on %s", cfg.GRPCPort)

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
