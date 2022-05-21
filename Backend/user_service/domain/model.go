package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Id             primitive.ObjectID `bson:"_id"`
	Username       string             `bson:"username"`
	HashedPassword string             `bson:"hashedPassword"`
	Role           string             `bson:"role"`
	FirstName      string             `bson:"firstName"`
	LastName       string             `bson:"lastName"`
	FullName       string             `bson:"fullName"`
	DateOfBirth    time.Time          `bson:"dateOfBirth"`
	Email          string             `bson:"email"`
	Activated      bool               `bson:"isActivated"`
	Private        bool               `bson:"isPrivate"`
	Token          string             `bson:"token"`
	PasswordToken  string             `bson:"passwordToken"`
}

func (user *User) IsCorrectPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	return err == nil
}
