package startup

import (
	"dislinkt/offer_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var offers = []*domain.Offer{}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
