package client

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	adminpb "txn-engine-phase-2/proto/gen/go/admin"
	walletpb "txn-engine-phase-2/proto/gen/go/wallet"
)

type AdminClient struct {
	conn         *grpc.ClientConn
	AuthClient   adminpb.AuthServiceClient
	WalletClient walletpb.WalletServiceClient
}

func NewAdminClient(addr string) *AdminClient {
	conn, err := grpc.Dial(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	return &AdminClient{
		conn:         conn,
		AuthClient:   adminpb.NewAuthServiceClient(conn),
		WalletClient: walletpb.NewWalletServiceClient(conn),
	}
}

func (c *AdminClient) Close() {
	if c.conn != nil {
		_ = c.conn.Close()
	}
}

func (c *AdminClient) Register(
	ctx context.Context,
	name, email, mobile, password, role string,
) (*adminpb.RegisterResponse, error) {

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return c.AuthClient.RegisterAdmin(ctx, &adminpb.RegisterRequest{
		Name:     name,
		Email:    email,
		Mobile:   mobile,
		Password: password,
		Role:     role,
	})
}

func (c *AdminClient) Login(
	ctx context.Context,
	mobile string,
) (*adminpb.LoginResponse, error) {

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return c.AuthClient.LoginAdmin(ctx, &adminpb.LoginRequest{
		Mobile: mobile,
	})
}

func (c *AdminClient) VerifyOtp(
	ctx context.Context,
	mobile, otp string,
) (*adminpb.VerifyOtpResponse, error) {

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return c.AuthClient.VerifyOtp(ctx, &adminpb.VerifyOtpRequest{
		Mobile: mobile,
		Otp:    otp,
	})
}

func (c *AdminClient) TopUpWallet(
	ctx context.Context,
	userID string,
	amount float64,
) (*walletpb.TopUpWalletResponse, error) {

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return c.WalletClient.TopUpWallet(ctx, &walletpb.TopUpWalletRequest{
		UserId: userID,
		Amount: amount,
	})
}
