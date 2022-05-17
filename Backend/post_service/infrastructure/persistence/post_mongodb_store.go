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
	DATABASE   = "post_service"
	COLLECTION = "post"
)

type PostMongoDBStore struct {
	posts *mongo.Collection
}

func NewPostMongoDBStore(client *mongo.Client) domain.PostStore {
	posts := client.Database(DATABASE).Collection(COLLECTION)
	return &PostMongoDBStore{
		posts: posts,
	}
}

func (store *PostMongoDBStore) Get(postId string) (*domain.Post, error) {
	id, err := primitive.ObjectIDFromHex(postId)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *PostMongoDBStore) GetAll() ([]*domain.Post, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *PostMongoDBStore) Create(post *domain.Post) error {
	result, err := store.posts.InsertOne(context.TODO(), post)
	if err != nil {
		return err
	}
	post.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *PostMongoDBStore) Update(postId string, post *domain.Post) error {
	id, err := primitive.ObjectIDFromHex(postId)
	if err != nil {
		return err
	}
	result, err := store.posts.ReplaceOne(
		context.TODO(),
		bson.M{"_id": id},
		post,
	)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New(post.Id.String())
	}
	return nil
}

func (store *PostMongoDBStore) DeleteAll() error {
	_, err := store.posts.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}

func (store *PostMongoDBStore) filter(filter interface{}) ([]*domain.Post, error) {
	cursor, err := store.posts.Find(context.TODO(), filter)
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *PostMongoDBStore) filterOne(filter interface{}) (post *domain.Post, err error) {
	result := store.posts.FindOne(context.TODO(), filter)
	err = result.Decode(&post)
	return
}

func decode(cursor *mongo.Cursor) (posts []*domain.Post, err error) {
	for cursor.Next(context.TODO()) {
		var Post domain.Post
		err = cursor.Decode(&Post)
		if err != nil {
			return
		}
		posts = append(posts, &Post)
	}
	err = cursor.Err()
	return
}
