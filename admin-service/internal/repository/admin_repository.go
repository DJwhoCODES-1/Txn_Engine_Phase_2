package repository

import (
	"context"
	"time"

	"txn-engine-phase-2/admin-service/internal/model"
	"txn-engine-phase-2/admin-service/internal/service"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AdminRepository struct {
	col *mongo.Collection
}

func NewAdminRepository(db *mongo.Database) *AdminRepository {
	return &AdminRepository{
		col: db.Collection("admins"),
	}
}

func (r *AdminRepository) Create(ctx context.Context, admin *model.Admin) (string, error) {
	res, err := r.col.InsertOne(ctx, admin)
	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r *AdminRepository) FindByMobile(ctx context.Context, mobile string) (*model.Admin, error) {
	var admin model.Admin
	err := r.col.FindOne(ctx, bson.M{"mobile": mobile}).Decode(&admin)
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *AdminRepository) UpdateOTP(ctx context.Context, mobile, otp string) error {
	_, err := r.col.UpdateOne(ctx,
		bson.M{"mobile": mobile},
		bson.M{
			"$set": bson.M{
				"otp":           otp,
				"otp_verified":  false,
				"otp_timestamp": time.Now(),
			},
		},
	)
	return err
}

func (r *AdminRepository) MarkOTPVerified(ctx context.Context, mobile string) error {
	_, err := r.col.UpdateOne(ctx,
		bson.M{"mobile": mobile},
		bson.M{
			"$set": bson.M{
				"otp_verified": true,
			},
		},
	)
	return err
}

func (r *AdminRepository) CreateTxn(txn *service.Transaction) error {
	_, err := r.col.InsertOne(context.Background(), txn)
	return err
}
