package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Post struct {
	Id        primitive.ObjectID `bson:"_id"`
	User      User               `bson:"user"`
	CreatedAt time.Time          `bson:"createdAt"`
	Content   Content            `bson:"content"`
}

type Content struct {
	Text  string `bson:"text"`
	Image string `bson:"image"`
	//Links []string `bson:"links"`
}

type User struct {
	Id        primitive.ObjectID `bson:"_id"`
	Username  string             `bson:"username"`
	FirstName string             `bson:"firstName"`
	LastName  string             `bson:"lastName"`
}

type Reaction struct {
	Id        primitive.ObjectID `bson:"_id"`
	UserId    primitive.ObjectID `bson:"userId"`
	PostId    primitive.ObjectID `bson:"postId"`
	Type      ReactionType       `bson:"type"`
	CreatedAt time.Time          `bson:"createdAt"`
}

type ReactionType int32

const (
	LIKE ReactionType = iota
	DISLIKE
)

type Comment struct {
	Id          primitive.ObjectID `bson:"_id"`
	Content     string             `bson:"content"`
	UserId      primitive.ObjectID `bson:"userId"`
	PostId      primitive.ObjectID `bson:"postId"`
	DateCreated time.Time          `bson:"dateCreated"`
}

type Connection struct {
	Id        primitive.ObjectID `bson:"_id"`
	IssuerId  primitive.ObjectID `bson:"_issuerId"`
	SubjectId primitive.ObjectID `bson:"_subjectId"`
}
