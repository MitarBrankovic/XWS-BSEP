package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Offer struct {
	Id          primitive.ObjectID `bson:"_id"`
	Username    string             `bson:"username"`
	Company     string             `bson:"company"`
	Description string             `bson:"description"`
	Position    string             `bson:"position"`
	Criteria    string             `bson:"criteria"`
	CreatedAt   time.Time          `bson:"createdAt"`
}
