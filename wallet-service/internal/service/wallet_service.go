package service

import (
	"context"
	"errors"
	"math"

	"txn-engine-phase-2/wallet-service/internal/repository"

	pb "txn-engine-phase-2/proto/gen/go/wallet"
)

type WalletService struct {
	pb.UnimplementedWalletServiceServer
	repo *repository.WalletRepository
}

func NewWalletService(repo *repository.WalletRepository) *WalletService {
	return &WalletService{repo: repo}
}

func (s *WalletService) TopUpWallet(ctx context.Context, req *pb.TopUpWalletRequest) (*pb.TopUpWalletResponse, error) {

	if req.UserId == "" {
		return nil, errors.New("userId required")
	}

	if req.Amount <= 0 {
		return nil, errors.New("invalid amount")
	}

	amount := math.Round(req.Amount*100) / 100

	wallet, err := s.repo.FindByUserID(req.UserId)
	if err != nil {
		return nil, err
	}

	prevBalance := wallet.Balance

	lienDeducted := math.Min(wallet.Lien, amount)

	wallet.Lien -= lienDeducted
	wallet.Balance += amount - lienDeducted
	wallet.Balance = math.Round(wallet.Balance*100) / 100

	if err := s.repo.Update(wallet); err != nil {
		return nil, err
	}

	return &pb.TopUpWalletResponse{
		PrevBalance:    prevBalance,
		UpdatedBalance: wallet.Balance,
		LienDeducted:   lienDeducted,
	}, nil
}
