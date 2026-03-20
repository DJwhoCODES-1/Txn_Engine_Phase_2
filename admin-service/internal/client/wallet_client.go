package client

import (
	"context"
	"time"

	pb "txn-engine-phase-2/proto/gen/go/wallet"

	"google.golang.org/grpc"
)

type WalletClient struct {
	client pb.WalletServiceClient
}

func NewWalletClient(conn *grpc.ClientConn) *WalletClient {
	return &WalletClient{
		client: pb.NewWalletServiceClient(conn),
	}
}

func (c *WalletClient) TopUp(ctx context.Context, userID string, amount float64) (*pb.TopUpWalletResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return c.client.TopUpWallet(ctx, &pb.TopUpWalletRequest{
		UserId: userID,
		Amount: amount,
	})
}
