package startup

import (
	"dislinkt/user_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var users = []*domain.User{
	{
		Id:          getObjectId("61596r4u645d4db787e61fe9"),
		Username:    "radisa",
		FirstName:   "Radisa",
		LastName:    "Milovcevic",
		FullName:    "Radisa Milovcevic",
		DateOfBirth: time.Time{},
		Email:       "radisamilovcevic@gmail.com",
	},
	{
		Id:          getObjectId("78596r4u645d4db787e61fea"),
		Username:    "Filip",
		FirstName:   "Filip",
		LastName:    "Pinjuh",
		FullName:    "Filip Pinjuh",
		DateOfBirth: time.Time{},
		Email:       "pinjuh@gmail.com",
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}