package startup

import (
	"dislinkt/user_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var users = []*domain.User{
	{
		Id:             getObjectId("61596r4u645d4db787e61fe9"),
		Username:       "radisa",
		HashedPassword: "$2a$10$KKYBMMbKxl4lSzDVbOxULOszwcSWZGW03O4A5gDhQlxadshhfiAc.",
		FirstName:      "Radisa",
		LastName:       "Milovcevic",
		FullName:       "Radisa Milovcevic",
		DateOfBirth:    time.Time{},
		Email:          "xwsclient@gmail.com",
		Role:           "user",
		Activated:      true,
		Private:        false,
		Skills:         []string{"aa", "bb", "cc"},
		Interests:      []string{"aa", "bb", "cc"},
		Education:      []domain.Education{{School: "osnovna", Degree: "6", FieldOfStudy: "stolar", StartDate: time.Time{}, EndDate: time.Time{}}},
		WorkExperience: []domain.WorkExperience{{Title: "Pomocnik stolara", Company: "Stolarske carolije", EmploymentType: "INTERNSHIP", Location: "Tutin", StartDate: time.Time{}, EndDate: time.Time{}}},
	},
	{
		Id:             getObjectId("78596r4u645d4db787e61fea"),
		Username:       "Filip",
		HashedPassword: "$2a$10$KKYBMMbKxl4lSzDVbOxULOszwcSWZGW03O4A5gDhQlxadshhfiAc.",
		FirstName:      "Filip",
		LastName:       "Pinjuh",
		FullName:       "Filip Pinjuh",
		DateOfBirth:    time.Time{},
		Email:          "pinjuh@gmail.com",
		Role:           "admin",
		Activated:      true,
		Private:        false,
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
