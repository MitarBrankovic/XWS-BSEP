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
	COLLECTIONREACTION = "reaction"
)

type ReactionMongoDBStore struct {
	reactions *mongo.Collection
}

func NewReactionMongoDBStore(client *mongo.Client) domain.ReactionStore {
	reactions := client.Database(DATABASE).Collection(COLLECTIONREACTION)
	return &ReactionMongoDBStore{
		reactions: reactions,
	}
}

func (store *ReactionMongoDBStore) Get(userId string) (*domain.Reaction, error) {
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *ReactionMongoDBStore) GetAll() ([]*domain.Reaction, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *ReactionMongoDBStore) Create(reaction *domain.Reaction) error {
	result, err := store.reactions.InsertOne(context.TODO(), reaction)
	if err != nil {
		return err
	}
	reaction.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *ReactionMongoDBStore) Update(reactionId string, reaction *domain.Reaction) error {
	id, err := primitive.ObjectIDFromHex(reactionId)
	if err != nil {
		return err
	}
	result, err := store.reactions.ReplaceOne(
		context.TODO(),
		bson.M{"_id": id},
		reaction,
	)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New(reaction.Id.String())
	}
	return nil
}

func (store *ReactionMongoDBStore) DeleteAll() error {
	_, err := store.reactions.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}

func (store *ReactionMongoDBStore) filter(filter interface{}) ([]*domain.Reaction, error) {
	cursor, err := store.reactions.Find(context.TODO(), filter)
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, context.TODO())

	if err != nil {
		return nil, err
	}
	return decodeReaction(cursor)
}

func (store *ReactionMongoDBStore) filterOne(filter interface{}) (reaction *domain.Reaction, err error) {
	result := store.reactions.FindOne(context.TODO(), filter)
	err = result.Decode(&reaction)
	return
}

func decodeReaction(cursor *mongo.Cursor) (reactions []*domain.Reaction, err error) {
	for cursor.Next(context.TODO()) {
		var Reaction domain.Reaction
		err = cursor.Decode(&Reaction)
		if err != nil {
			return
		}
		reactions = append(reactions, &Reaction)
	}
	err = cursor.Err()
	return
}
