package persistence

import (
	"context"
	"dislinkt/offer_service/domain"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "offer_service"
	COLLECTION = "offer"
)

type OfferMongoDBStore struct {
	offers *mongo.Collection
}

func NewOfferMongoDBStore(client *mongo.Client) domain.OfferStore {
	offers := client.Database(DATABASE).Collection(COLLECTION)
	return &OfferMongoDBStore{
		offers: offers,
	}
}

func (store *OfferMongoDBStore) Get(offerId string) (*domain.Offer, error) {
	id, err := primitive.ObjectIDFromHex(offerId)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *OfferMongoDBStore) GetAll() ([]*domain.Offer, error) {
	filter := bson.M{"role": "offer"}
	return store.filter(filter)
}

func (store *OfferMongoDBStore) Create(offer *domain.Offer) error {
	offer.Id = primitive.NewObjectID()
	result, err := store.offers.InsertOne(context.TODO(), offer)
	if err != nil {
		return err
	}
	offer.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *OfferMongoDBStore) Update(offerId string, offer *domain.Offer) error {
	id, err := primitive.ObjectIDFromHex(offerId)
	if err != nil {
		return err
	}
	result, err := store.offers.ReplaceOne(
		context.TODO(),
		bson.M{"_id": id},
		offer,
	)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New(offer.Id.String())
	}
	return nil
}

func (store *OfferMongoDBStore) DeleteAll() error {
	_, err := store.offers.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}

func (store *OfferMongoDBStore) filter(filter interface{}) ([]*domain.Offer, error) {
	cursor, err := store.offers.Find(context.TODO(), filter)
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

func (store *OfferMongoDBStore) filterOne(filter interface{}) (offer *domain.Offer, err error) {
	result := store.offers.FindOne(context.TODO(), filter)
	err = result.Decode(&offer)
	return
}

func decode(cursor *mongo.Cursor) (offers []*domain.Offer, err error) {
	for cursor.Next(context.TODO()) {
		var Offer domain.Offer
		err = cursor.Decode(&Offer)
		if err != nil {
			return
		}
		offers = append(offers, &Offer)
	}
	err = cursor.Err()
	return
}
