package server

import (
	"net"

	"txn-engine-phase-2/wallet-service/internal/config"
	"txn-engine-phase-2/wallet-service/internal/repository"
	"txn-engine-phase-2/wallet-service/internal/service"

	pb "txn-engine-phase-2/proto/gen/go/wallet"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	cfg *config.Config
	db  *mongo.Database
}

func NewGRPCServer(cfg *config.Config, db *mongo.Database) *GRPCServer {
	return &GRPCServer{cfg: cfg, db: db}
}

func (s *GRPCServer) Run() error {
	addr := s.cfg.GRPCPort
	if addr[0] != ':' {
		addr = ":" + addr
	}

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	repo := repository.NewWalletRepository(s.db)
	svc := service.NewWalletService(repo)

	grpcServer := grpc.NewServer()

	pb.RegisterWalletServiceServer(grpcServer, svc)

	return grpcServer.Serve(lis)
}
