package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Admin struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name         string             `bson:"name"`
	Email        string             `bson:"email"`
	Mobile       string             `bson:"mobile"`
	Password     string             `bson:"password"`
	Role         string             `bson:"role"`
	OTP          string             `bson:"otp,omitempty"`
	OTPVerified  bool               `bson:"otp_verified"`
	OTPTimestamp time.Time          `bson:"otp_timestamp,omitempty"`
	CreatedAt    time.Time          `bson:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt"`
}
