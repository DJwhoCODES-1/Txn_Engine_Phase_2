package main

import (
	"log"
	"net"

	"txn-engine-phase-2/admin-service/internal/client"
	"txn-engine-phase-2/admin-service/internal/config"
	"txn-engine-phase-2/admin-service/internal/database"
	"txn-engine-phase-2/admin-service/internal/repository"
	"txn-engine-phase-2/admin-service/internal/server"
	"txn-engine-phase-2/admin-service/internal/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cfg := config.Load()

	db := database.ConnectMongo(cfg.MongoURI)

	// gRPC connection to wallet-service
	conn, err := grpc.Dial(cfg.WalletServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	walletClient := client.NewWalletClient(conn)

	txnRepo := repository.NewTxnRepository(db)

	walletService := service.NewWalletService(walletClient, txnRepo)

	grpcServer := server.NewGRPCServer(walletService)

	lis, err := net.Listen("tcp", ":"+cfg.GRPCPort)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Admin service running on %s", cfg.GRPCPort)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
