package repository

import (
	"context"
	"errors"
	"time"

	"txn-engine-phase-2/wallet-service/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type WalletRepository struct {
	col *mongo.Collection
}

func NewWalletRepository(db *mongo.Database) *WalletRepository {
	return &WalletRepository{
		col: db.Collection("wallets"),
	}
}

func (r *WalletRepository) FindByUserID(userID string) (*model.Wallet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var wallet model.Wallet
	err := r.col.FindOne(ctx, bson.M{"userId": userID}).Decode(&wallet)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("wallet not found")
		}
		return nil, err
	}
	return &wallet, nil
}

func (r *WalletRepository) Update(wallet *model.Wallet) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.col.UpdateOne(
		ctx,
		bson.M{"userId": wallet.UserID},
		bson.M{
			"$set": bson.M{
				"balance": wallet.Balance,
				"lien":    wallet.Lien,
			},
		},
	)

	return err
}
