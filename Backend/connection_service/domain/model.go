package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Connection struct {
	Id          primitive.ObjectID `bson:"_id"`
	IssuerUser  User               `bson:"issuerUser"`
	SubjectUser User               `bson:"subjectUser"`
	Date        time.Time          `bson:"date"`
	IsApproved  bool               `bson:"isApproved"`
}

type Message struct {
	Id               primitive.ObjectID `bson:"_id"`
	SenderUsername   string             `bson:"senderUsername"`
	ReceiverUsername string             `bson:"receiverUsername"`
	Date             time.Time          `bson:"date"`
	Content          string             `bson:"content"`
}

type User struct {
	Username string `bson:"username"`
	Private  bool   `bson:"private"`
}
