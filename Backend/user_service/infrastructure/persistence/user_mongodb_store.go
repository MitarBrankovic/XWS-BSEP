package persistence

import (
	"context"
	"dislinkt/user_service/domain"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "user_service"
	COLLECTION = "user"
)

type UserMongoDBStore struct {
	users *mongo.Collection
}

func NewUserMongoDBStore(client *mongo.Client) domain.UserStore {
	users := client.Database(DATABASE).Collection(COLLECTION)
	return &UserMongoDBStore{
		users: users,
	}
}

func (store *UserMongoDBStore) Get(userId string) (*domain.User, error) {
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *UserMongoDBStore) GetAll() ([]*domain.User, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *UserMongoDBStore) Create(user *domain.User) error {
	user.Id = primitive.NewObjectID()
	result, err := store.users.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	user.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *UserMongoDBStore) Update(userId string, user *domain.User) error {
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}
	result, err := store.users.ReplaceOne(
		context.TODO(),
		bson.M{"_id": id},
		user,
	)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New(user.Id.String())
	}
	return nil
}

func (store *UserMongoDBStore) DeleteAll() error {
	_, err := store.users.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}

func (store *UserMongoDBStore) Find(username string) (*domain.User, error) {
	filter := bson.M{"username": username}
	return store.filterOne(filter)
}

func (store *UserMongoDBStore) ActivateAccount(token string) *domain.User {
	filter := bson.M{"token": token}
	user, err := store.filterOne(filter)
	user.Activated = true
	user.Token = ""
	_, err = store.users.ReplaceOne(
		context.TODO(),
		bson.M{"_id": user.Id},
		user,
	)
	if err != nil {
		return user
	}
	return user
}

func (store *UserMongoDBStore) PasswordlessLoginDemand(username string) (*domain.User, error) {
	filter := bson.M{"username": username}
	return store.filterOne(filter)
}

func (store *UserMongoDBStore) PasswordlessLogin(token string) (*domain.User, error) {
	filter := bson.M{"passwordToken": token}
	return store.filterOne(filter)
}

func (store *UserMongoDBStore) filter(filter interface{}) ([]*domain.User, error) {
	cursor, err := store.users.Find(context.TODO(), filter)
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

func (store *UserMongoDBStore) filterOne(filter interface{}) (user *domain.User, err error) {
	result := store.users.FindOne(context.TODO(), filter)
	err = result.Decode(&user)
	return
}

func decode(cursor *mongo.Cursor) (users []*domain.User, err error) {
	for cursor.Next(context.TODO()) {
		var User domain.User
		err = cursor.Decode(&User)
		if err != nil {
			return
		}
		users = append(users, &User)
	}
	err = cursor.Err()
	return
}
