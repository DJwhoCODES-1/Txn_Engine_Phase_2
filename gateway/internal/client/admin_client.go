package client

import (
	"context"
	"log"
	"time"

	adminpb "txn-engine-phase-2/proto/gen/go/admin"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AdminClient struct {
	Client adminpb.AuthServiceClient
}

func NewAdminClient(addr string) *AdminClient {
	conn, err := grpc.Dial(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect to admin service: %v", err)
	}

	return &AdminClient{
		Client: adminpb.NewAuthServiceClient(conn),
	}
}

func (c *AdminClient) Login(ctx context.Context, mobile string) (*adminpb.LoginResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return c.Client.LoginAdmin(ctx, &adminpb.LoginRequest{
		Mobile: mobile,
	})
}

func (c *AdminClient) VerifyOtp(ctx context.Context, mobile, otp string) (*adminpb.VerifyOtpResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return c.Client.VerifyOtp(ctx, &adminpb.VerifyOtpRequest{
		Mobile: mobile,
		Otp:    otp,
	})
}
