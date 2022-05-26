package startup

import (
	"dislinkt/post_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var posts = []*domain.Post{
	{
		Id: getObjectId("11596r4u645d4db787e61fe9"),
		User: domain.User{
			Id:        getObjectId("61596r4u645d4db787e61fe9"),
			FirstName: "radisa",
			LastName:  "Radisa",
		},
		CreatedAt: time.Time{},
		Content: domain.Content{
			Text:  "kontent",
			Image: "slika",
		},
	},
}

var reactions = []*domain.Reaction{
	{
		Id:        getObjectId("21596r4u645d4db787e61fe9"),
		UserId:    getObjectId("61596r4u645d4db787e61fe9"),
		PostId:    getObjectId("11596r4u645d4db787e61fe9"),
		Type:      0,
		CreatedAt: time.Time{},
	},
}

var comments = []*domain.Comment{
	{
		Id:          getObjectId("31596r4u645d4db787e61fe9"),
		Content:     "kontent",
		UserId:      getObjectId("61596r4u645d4db787e61fe9"),
		PostId:      getObjectId("11596r4u645d4db787e61fe9"),
		DateCreated: time.Time{},
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}

/*domain.User{
Id:        getObjectId("61596r4u645d4db787e61fe9"),
FirstName: "radisa",
LastName:  "Radisa",
}*/
