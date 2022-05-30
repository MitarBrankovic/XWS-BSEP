package persistence

import (
	"context"
	"dislinkt/connection_service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	DATABASE    = "connection_service"
	COLLECTION1 = "connection"
)

type ConnectionMongoDBStore struct {
	connections *mongo.Collection
}

func NewConnectionMongoDBStore(client *mongo.Client) domain.ConnectionStore {
	connections := client.Database(DATABASE).Collection(COLLECTION1)
	return &ConnectionMongoDBStore{
		connections: connections,
	}
}

func (store *ConnectionMongoDBStore) Get(userId string) ([]*domain.Connection, error) {
	filter := bson.M{"$or": []bson.M{{"subjectId": userId},
		{"issuerId": userId}}}
	return store.filter(filter)
}

func (store *ConnectionMongoDBStore) GetAll() ([]*domain.Connection, error) {
	filter := bson.M{}
	return store.filter(filter)
}

func (store *ConnectionMongoDBStore) Create(connection *domain.Connection) (*domain.Connection, error) {
	connection.IsApproved = false
	if connection.IssuerUsername == connection.SubjectUsername {
		return nil, status.Errorf(codes.InvalidArgument, "Cannot create connection with same username")
	}

	filter := bson.D{{}}
	allConnections, _ := store.filter(filter)
	for _, c := range allConnections {
		if c.IssuerUsername == connection.IssuerUsername && c.SubjectUsername == connection.SubjectUsername {
			return nil, status.Errorf(codes.AlreadyExists, "Connection already exists")
		}
	}
	result, err := store.connections.InsertOne(context.TODO(), connection)
	if err != nil {
		return nil, err
	}
	connection.Id = result.InsertedID.(primitive.ObjectID)
	return connection, nil
}

func (store *ConnectionMongoDBStore) DeleteAll() error {
	_, err := store.connections.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}

func (store *ConnectionMongoDBStore) Delete(id string) error {
	Id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": Id}
	_, err = store.connections.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (store *ConnectionMongoDBStore) Update(id string) (*domain.Connection, error) {
	Id, err := primitive.ObjectIDFromHex(id)
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
	}
	return connection, nil
}

func (store *ConnectionMongoDBStore) filter(filter interface{}) ([]*domain.Connection, error) {
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
	return decode(cursor)
}

func (store *ConnectionMongoDBStore) filterOne(filter interface{}) (connection *domain.Connection, err error) {
	result := store.connections.FindOne(context.TODO(), filter)
	err = result.Decode(&connection)
	return
}

func decode(cursor *mongo.Cursor) (connections []*domain.Connection, err error) {
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
