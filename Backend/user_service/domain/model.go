package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	Id          primitive.ObjectID `bson:"_id"`
	Username    string             `bson:"username"`
	FirstName   string             `bson:"firstName"`
	LastName    string             `bson:"lastName"`
	FullName    string             `bson:"fullName"`
	DateOfBirth time.Time          `bson:"dateOfBirth"`
	Email       string             `bson:"email"`
}
