package startup

import (
	"dislinkt/connection_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var connections = []*domain.Connection{
	{
		Id:         getObjectId("45406d1b624b3da748f63fe1"),
		IssuerId:   getObjectId("61596r4u645d4db787e61fe9"),
		SubjectId:  getObjectId("78596r4u645d4db787e61fea"),
		IsApproved: true,
		Date:       time.Now(),
	},
}

var profilesPrivacy = []*domain.ProfilePrivacy{
	{
		Id:        primitive.NewObjectID(),
		UserId:    getObjectId("61596r4u645d4db787e61fe9"),
		IsPrivate: false,
	},
	{
		Id:        primitive.NewObjectID(),
		UserId:    getObjectId("78596r4u645d4db787e61fea"),
		IsPrivate: false,
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
