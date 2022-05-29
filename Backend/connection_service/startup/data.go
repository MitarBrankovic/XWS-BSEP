package startup

import (
	"dislinkt/connection_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var connections = []*domain.Connection{
	{
		Id:              getObjectId("45406d1b624b3da748f63fe1"),
		IssuerUsername:  "radisa",
		SubjectUsername: "filip",
		IsApproved:      true,
		Date:            time.Now(),
	},
}

var messages = []*domain.Message{
	{
		Id:               getObjectId("45406d1b624b3da748f63fe1"),
		SenderUsername:   "radisa",
		ReceiverUsername: "filip",
		Content:          "Hello",
		Date:             time.Now(),
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
