package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Connection struct {
	Id              primitive.ObjectID `bson:"_id"`
	IssuerUsername  string             `bson:"issuerUsername"`
	SubjectUsername string             `bson:"subjectUsername"`
	Date            time.Time          `bson:"date"`
	IsApproved      bool               `bson:"isApproved"`
}
