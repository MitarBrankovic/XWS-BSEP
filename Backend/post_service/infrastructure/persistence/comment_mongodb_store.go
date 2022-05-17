package persistence

import (
	"context"
	"dislinkt/post_service/domain"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	COLLECTIONCOMMENT = "comment"
)

type CommentMongoDBStore struct {
	comments *mongo.Collection
}

func NewCommentMongoDBStore(client *mongo.Client) domain.CommentStore {
	comments := client.Database(DATABASE).Collection(COLLECTIONCOMMENT)
	return &CommentMongoDBStore{
		comments: comments,
	}
}

func (store *CommentMongoDBStore) Get(commentId string) (*domain.Comment, error) {
	id, err := primitive.ObjectIDFromHex(commentId)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *CommentMongoDBStore) GetAll() ([]*domain.Comment, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *CommentMongoDBStore) Create(comment *domain.Comment) error {
	result, err := store.comments.InsertOne(context.TODO(), comment)
	if err != nil {
		return err
	}
	comment.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *CommentMongoDBStore) Update(commentId string, comment *domain.Comment) error {
	id, err := primitive.ObjectIDFromHex(commentId)
	if err != nil {
		return err
	}
	result, err := store.comments.ReplaceOne(
		context.TODO(),
		bson.M{"_id": id},
		comment,
	)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New(comment.Id.String())
	}
	return nil
}

func (store *CommentMongoDBStore) DeleteAll() error {
	_, err := store.comments.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}

func (store *CommentMongoDBStore) filter(filter interface{}) ([]*domain.Comment, error) {
	cursor, err := store.comments.Find(context.TODO(), filter)
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, context.TODO())

	if err != nil {
		return nil, err
	}
	return decodeComment(cursor)
}

func (store *CommentMongoDBStore) filterOne(filter interface{}) (comment *domain.Comment, err error) {
	result := store.comments.FindOne(context.TODO(), filter)
	err = result.Decode(&comment)
	return
}

func decodeComment(cursor *mongo.Cursor) (comments []*domain.Comment, err error) {
	for cursor.Next(context.TODO()) {
		var Comment domain.Comment
		err = cursor.Decode(&Comment)
		if err != nil {
			return
		}
		comments = append(comments, &Comment)
	}
	err = cursor.Err()
	return
}
