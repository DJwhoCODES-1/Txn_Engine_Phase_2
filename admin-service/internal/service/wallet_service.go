package service

import (
	"context"
	"errors"

	"txn-engine-phase-2/admin-service/internal/client"
	"txn-engine-phase-2/admin-service/internal/utils"
)

type WalletService struct {
	walletClient *client.WalletClient
	txnRepo      TxnRepository
}

type TxnRepository interface {
	Create(txn *Transaction) error
}

type Transaction struct {
	TransactionID  string
	ClientReqID    string
	UserID         string
	Amount         float64
	PrevBalance    float64
	UpdatedBalance float64
	Status         string
}

func NewWalletService(wc *client.WalletClient, tr TxnRepository) *WalletService {
	return &WalletService{
		walletClient: wc,
		txnRepo:      tr,
	}
}

func (s *WalletService) TopUpWallet(ctx context.Context, merchant map[string]interface{}, amount float64, admin map[string]interface{}) (*Transaction, error) {

	userID, ok := merchant["userId"].(string)
	if !ok || userID == "" {
		return nil, errors.New("userId required")
	}

	resp, err := s.walletClient.TopUp(ctx, userID, amount)
	if err != nil {
		return nil, err
	}

	txn := &Transaction{
		TransactionID:  utils.GenerateReqID(),
		ClientReqID:    utils.GenerateReqID(),
		UserID:         admin["userId"].(string),
		Amount:         amount,
		PrevBalance:    resp.PrevBalance,
		UpdatedBalance: resp.UpdatedBalance,
		Status:         "success",
	}

	err = s.txnRepo.Create(txn)
	if err != nil {
		return nil, err
	}

	return txn, nil
}
