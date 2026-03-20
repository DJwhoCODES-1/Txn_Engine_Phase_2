package main

import (
	"log"

	"txn-engine-phase-2/admin-service/internal/config"
	"txn-engine-phase-2/admin-service/internal/database"
	"txn-engine-phase-2/admin-service/internal/server"
)

func main() {
	cfg := config.Load()

	db := database.ConnectMongo(cfg.MongoURI)

	s := server.NewGRPCServer(cfg, db)

	log.Printf("Admin service running on %s", cfg.GRPCPort)

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
