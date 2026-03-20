package server

import (
	"net"

	"txn-engine-phase-2/admin-service/internal/config"
	"txn-engine-phase-2/admin-service/internal/handler"
	"txn-engine-phase-2/admin-service/internal/repository"
	"txn-engine-phase-2/admin-service/internal/service"
	adminpb "txn-engine-phase-2/proto/gen/go/admin"

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

	repo := repository.NewAdminRepository(s.db)
	svc := service.NewAuthService(repo, s.cfg.JWTSecret)
	handler := handler.NewGRPCHandler(svc)

	grpcServer := grpc.NewServer()

	adminpb.RegisterAuthServiceServer(grpcServer, handler)

	return grpcServer.Serve(lis)
}
