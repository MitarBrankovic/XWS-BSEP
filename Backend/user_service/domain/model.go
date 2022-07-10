package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Id                  primitive.ObjectID `bson:"_id"`
	Username            string             `bson:"username" validate:"username"`
	HashedPassword      string             `bson:"hashedPassword"`
	Role                string             `bson:"role"`
	FirstName           string             `bson:"firstName"`
	LastName            string             `bson:"lastName"`
	FullName            string             `bson:"fullName" validate:"name"`
	DateOfBirth         time.Time          `bson:"dateOfBirth"`
	Email               string             `bson:"email" validate:"email"`
	Activated           bool               `bson:"activated"`
	Private             bool               `bson:"private"`
	Token               string             `bson:"token"`
	TokenDate           time.Time          `bson:"tokenDate"`
	PasswordToken       string             `bson:"passwordToken"`
	PasswordTokenDate   time.Time          `bson:"passwordTokenDate"`
	RecoveryToken       string             `bson:"recoveryToken"`
	RecoveryTokenDate   time.Time          `bson:"recoveryTokenDate"`
	TwoFactorToken      string             `bson:"twoFactorToken"`
	TwoFactorTokenDate  time.Time          `bson:"twoFactorTokenDate"`
	TwoFactorEnabled    bool               `bson:"twoFactorEnabled"`
	Education           []Education        `bson:"education"`
	WorkExperience      []WorkExperience   `bson:"workExperience"`
	Skills              []string           `bson:"skills"`
	Interests           []string           `bson:"interests"`
	ApiToken            string             `bson:"apiToken"`
	FollowNotification  bool               `bson:"followNotification"`
	PostNotification    bool               `bson:"postNotification"`
	MessageNotification bool               `bson:"messageNotification"`
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

type Block struct {
	Id              primitive.ObjectID `bson:"_id"`
	IssuerUsername  string             `bson:"issuerUsername"`
	SubjectUsername string             `bson:"subjectUsername"`
}

type NotificationType int32

const (
	FOLLOW NotificationType = iota
	POST
	MESSAGE
)

type Notification struct {
	Id       primitive.ObjectID `bson:"_id"`
	Username string             `bson:"username"`
	Type     NotificationType   `bson:"type"`
	Message  string             `bson:"message"`
}

func (user *User) IsCorrectPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	return err == nil
}
