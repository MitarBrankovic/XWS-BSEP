package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Id             primitive.ObjectID `bson:"_id"`
	Username       string             `bson:"username" validate:"username"`
	HashedPassword string             `bson:"hashedPassword"`
	Role           string             `bson:"role"`
	FirstName      string             `bson:"firstName"`
	LastName       string             `bson:"lastName"`
	FullName       string             `bson:"fullName" validate:"name"`
	DateOfBirth    time.Time          `bson:"dateOfBirth"`
	Email          string             `bson:"email" validate:"email"`
	Activated      bool               `bson:"isActivated"`
	Private        bool               `bson:"isPrivate"`
	Token          string             `bson:"token"`
	PasswordToken  string             `bson:"passwordToken"`
	RecoveryToken  string             `bson:"recoveryToken"`
	Education      []Education        `bson:"education"`
	WorkExperience []WorkExperience   `bson:"workExperience"`
	Skills         []string           `bson:"skills"`
	Interests      []string           `bson:"interests"`
}

type Education struct {
	School       string    `bson:"school"`
	Degree       string    `bson:"degree"`
	FieldOfStudy string    `bson:"fieldOfStudy"`
	StartDate    time.Time `bson:"startDate"`
	EndDate      time.Time `bson:"endDate"`
}

type WorkExperience struct {
	Title          string    `bson:"title"`
	Company        string    `bson:"company"`
	EmploymentType string    `bson:"employmentType"`
	Location       string    `bson:"location"`
	StartDate      time.Time `bson:"startDate"`
	EndDate        time.Time `bson:"endDate"`
}

func (user *User) IsCorrectPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	return err == nil
}
