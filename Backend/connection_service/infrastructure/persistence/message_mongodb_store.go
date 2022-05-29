package persistence

import (
	"context"
	"dislinkt/connection_service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	COLLECTION = "message"
)

type MessageMongoDBStore struct {
	messages *mongo.Collection
}

func NewMessageMongoDBStore(client *mongo.Client) domain.MessageStore {
	messages := client.Database(DATABASE).Collection(COLLECTION)
	return &MessageMongoDBStore{
		messages: messages,
	}
}

func (store *MessageMongoDBStore) Get(userId string) ([]*domain.Message, error) {
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"$or": []bson.M{{"_id": id},
		{"_id": id}}}
	return store.filter(filter)
}

func (store *MessageMongoDBStore) Create(message *domain.Message) (*domain.Message, error) {
	message.Id = primitive.NewObjectID()
	result, err := store.messages.InsertOne(context.TODO(), message)
	if err != nil {
		return nil, err
	}
	message.Id = result.InsertedID.(primitive.ObjectID)
	return message, nil
}

func (store *MessageMongoDBStore) DeleteAll() error {
	_, err := store.messages.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}

func (store *MessageMongoDBStore) Delete(id string) error {
	Id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": Id}
	_, err = store.messages.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (store *MessageMongoDBStore) Update(id string) (*domain.Message, error) {
	/*Id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": Id}
	connection, err := store.filterOne(filter)
	if err != nil {
		return nil, err
	}
	connection.IsApproved = true
	_, err = store.connections.UpdateOne(context.TODO(), filter, bson.D{{"$set", bson.M{"isApproved": connection.IsApproved}}})
	if err != nil {
		return nil, err
	}*/
	return nil, nil
}

func (store *MessageMongoDBStore) filter(filter interface{}) ([]*domain.Message, error) {
	cursor, err := store.messages.Find(context.TODO(), filter)
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, context.TODO())

	if err != nil {
		return nil, err
	}
	return decodeMessage(cursor)
}

func (store *MessageMongoDBStore) filterOne(filter interface{}) (message *domain.Message, err error) {
	result := store.messages.FindOne(context.TODO(), filter)
	err = result.Decode(&message)
	return
}

func decodeMessage(cursor *mongo.Cursor) (messages []*domain.Message, err error) {
	for cursor.Next(context.TODO()) {
		var Message domain.Message
		err = cursor.Decode(&Message)
		if err != nil {
			return
		}
		messages = append(messages, &Message)
	}
	err = cursor.Err()
	return
}
