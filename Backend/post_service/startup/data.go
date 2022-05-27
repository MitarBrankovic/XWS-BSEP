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
			Username:  "radisa",
			FirstName: "Radisa",
			LastName:  "Milovcevic",
		},
		CreatedAt: time.Time{},
		Content: domain.Content{
			Text:  "kontent",
			Image: "slika",
		},
		Comments: []domain.Comment{
			{
				Id: getObjectId("11596r4u645d4db787e61fe8"),
				User: domain.User{
					Id:        getObjectId("61596r4u645d4db787e61fe9"),
					Username:  "radisa",
					FirstName: "Radisa",
					LastName:  "Milovcevic",
				},
				DateCreated: time.Time{},
				Content:     "neki komentar",
			},
		},
		Reactions: []domain.Reaction{
			{
				Id: getObjectId("11596r4u645d4db787e61fe7"),
				User: domain.User{
					Id:        getObjectId("61596r4u645d4db787e61fe9"),
					Username:  "radisa",
					FirstName: "Radisa",
					LastName:  "Milovcevic",
				},
				CreatedAt: time.Time{},
				Type:      domain.ReactionType(1),
			},
		},
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
