package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Wallet struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	UserID  string             `bson:"userId"`
	Balance float64            `bson:"balance"`
	Lien    float64            `bson:"lien"`
}
