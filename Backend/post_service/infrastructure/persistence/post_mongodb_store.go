package persistence

import (
	"context"
	pb "dislinkt/common/proto/post_service"
	"dislinkt/post_service/domain"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE    = "post_service"
	COLLECTION  = "post"
	COLLECTION2 = "connection"
)

type PostMongoDBStore struct {
	posts       *mongo.Collection
	connections *mongo.Collection
}

func NewPostMongoDBStore(client *mongo.Client) domain.PostStore {
	posts := client.Database(DATABASE).Collection(COLLECTION)
	connections := client.Database(DATABASE).Collection(COLLECTION2)
	return &PostMongoDBStore{
		posts:       posts,
		connections: connections,
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
	post.Id = primitive.NewObjectID()
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

func (store *PostMongoDBStore) UpdateUser(username string, user *pb.User) (*domain.User, error) {
	filter := bson.D{{}}
	posts, _ := store.filter(filter)
	changedUser := domain.User{}
	for _, post := range posts {
		if post.User.Username == username {
			post.User.FirstName = user.FirstName
			post.User.LastName = user.LastName
			changedUser = post.User
			_, err := store.posts.ReplaceOne(
				context.TODO(),
				bson.M{"_id": post.Id},
				post,
			)
			if err != nil {
				return nil, err
			}
		}
	}
	return &changedUser, nil
}

func (store *PostMongoDBStore) GetByUser(username string) ([]*domain.Post, error) {
	filter := bson.D{{}}
	posts, _ := store.filter(filter)

	userPosts := []*domain.Post{}

	for _, post := range posts {
		if post.User.Username == username {
			userPosts = append(userPosts, post)
		}
	}

	return userPosts, nil
}

func (store *PostMongoDBStore) AddReaction(reaction *domain.Reaction, postId string) error {
	reaction.Id = primitive.NewObjectID()
	id, err := primitive.ObjectIDFromHex(postId)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": id}
	post, err := store.filterOne(filter)
	if err != nil {
		return err
	}
	post.Reactions = append(post.Reactions, *reaction)
	_, err = store.posts.ReplaceOne(
		context.TODO(),
		bson.M{"_id": post.Id},
		post,
	)
	if err != nil {
		return err
	}
	return nil
}

func (store *PostMongoDBStore) AddComment(comment *domain.Comment, postId string) error {
	comment.Id = primitive.NewObjectID()
	id, err := primitive.ObjectIDFromHex(postId)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": id}
	post, err := store.filterOne(filter)
	if err != nil {
		return err
	}
	post.Comments = append(post.Comments, *comment)
	_, err = store.posts.ReplaceOne(
		context.TODO(),
		bson.M{"_id": post.Id},
		post,
	)
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

func (store *PostMongoDBStore) GetProfilePosts(profileId string) ([]*domain.Post, error) {
	id, err := primitive.ObjectIDFromHex(profileId)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"profile._id": id}
	return store.filter(filter)
}

func (store *PostMongoDBStore) GetConnectionPosts(profileId string) ([]*domain.Post, error) {
	id, err := primitive.ObjectIDFromHex(profileId)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{"$or", bson.A{bson.M{"_issuerId": id}, bson.M{"_subjectId": id}}}}
	connections, err := store.filterConnections(filter)
	if err != nil {
		return nil, err
	}
	posts := make([]*domain.Post, 0)
	for _, connection := range connections {
		if connection.IssuerId == id {
			connectionPosts, err := store.filter(bson.M{"profile._id": connection.SubjectId})
			if err != nil {
				return nil, err
			}
			posts = append(posts, connectionPosts...)
		} else if connection.SubjectId == id {
			connectionPosts, err := store.filter(bson.M{"profile._id": connection.IssuerId})
			if err != nil {
				return nil, err
			}
			posts = append(posts, connectionPosts...)
		}
	}
	return posts, nil
}

func (store *PostMongoDBStore) filterConnections(filter interface{}) ([]*domain.Connection, error) {
	cursor, err := store.connections.Find(context.TODO(), filter)
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, context.TODO())

	if err != nil {
		return nil, err
	}
	return decodeConnections(cursor)
}

func decodeConnections(cursor *mongo.Cursor) (connections []*domain.Connection, err error) {
	for cursor.Next(context.TODO()) {
		var Connection domain.Connection
		err = cursor.Decode(&Connection)
		if err != nil {
			return
		}
		connections = append(connections, &Connection)
	}
	err = cursor.Err()
	return
}
