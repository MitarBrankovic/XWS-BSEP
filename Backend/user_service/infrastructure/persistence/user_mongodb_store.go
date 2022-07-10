package persistence

import (
	"context"
	"dislinkt/user_service/domain"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

const (
	DATABASE               = "user_service"
	COLLECTION             = "user"
	BLOCKCOLLECTION        = "block"
	NOTIFICATIONCOLLECTION = "notification"
)

type UserMongoDBStore struct {
	users         *mongo.Collection
	blocks        *mongo.Collection
	notifications *mongo.Collection
}

func NewUserMongoDBStore(client *mongo.Client) domain.UserStore {
	users := client.Database(DATABASE).Collection(COLLECTION)
	blocks := client.Database(DATABASE).Collection(BLOCKCOLLECTION)
	notifications := client.Database(DATABASE).Collection(NOTIFICATIONCOLLECTION)
	return &UserMongoDBStore{
		users:         users,
		blocks:        blocks,
		notifications: notifications,
	}
}

func (store *UserMongoDBStore) Block(block *domain.Block) error {
	block.Id = primitive.NewObjectID()
	filter := bson.M{}
	blocks, err := store.filterBlock(filter)
	for _, b := range blocks {
		if b.IssuerUsername == block.IssuerUsername && b.SubjectUsername == block.SubjectUsername {
			return errors.New("block already exists")
		}
	}
	result, err := store.blocks.InsertOne(context.TODO(), block)
	if err != nil {
		return err
	}
	block.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *UserMongoDBStore) UnBlock(block *domain.Block) error {
	filter := bson.M{"issuerUsername": block.IssuerUsername, "subjectUsername": block.SubjectUsername}
	block, err := store.filterOneBlock(filter)
	if err != nil {
		return err
	}
	if block == nil {
		return errors.New("block not found")
	}
	_, err = store.blocks.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (store *UserMongoDBStore) CreateNotification(notification *domain.Notification) error {
	notification.Id = primitive.NewObjectID()
	result, err := store.notifications.InsertOne(context.TODO(), notification)
	if err != nil {
		return err
	}
	notification.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *UserMongoDBStore) DeleteNotification(notification *domain.Notification) error {
	filter := bson.M{"_id": notification.Id}
	_, err := store.notifications.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (store *UserMongoDBStore) GetNotifications(username string) ([]*domain.Notification, error) {
	filter := bson.M{"username": username}
	return store.filterNotifications(filter)
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
	filter := bson.M{"role": "user"}
	return store.filter(filter)
}

func (store *UserMongoDBStore) GetAllBlock() ([]*domain.Block, error) {
	filter := bson.M{}
	return store.filterBlock(filter)
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

func (store *UserMongoDBStore) DeleteAllBlocks() error {
	_, err := store.blocks.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}

func (store *UserMongoDBStore) DeleteAllNotifications() error {
	_, err := store.notifications.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}
	return nil
}

func (store *UserMongoDBStore) Find(username string) (*domain.User, error) {
	filter := bson.M{"username": username}
	return store.filterOne(filter)
}

func (store *UserMongoDBStore) FindByEmail(email string) (*domain.User, error) {
	filter := bson.M{"email": email}
	return store.filterOne(filter)
}

func (store *UserMongoDBStore) FindByActivationToken(token string) (*domain.User, error) {
	filter := bson.M{"token": token}
	return store.filterOne(filter)
}

func (store *UserMongoDBStore) FindByPasswordlessToken(token string) (*domain.User, error) {
	filter := bson.M{"passwordToken": token}
	return store.filterOne(filter)
}

func (store *UserMongoDBStore) FindByRecoveryToken(token string) (*domain.User, error) {
	filter := bson.M{"recoveryToken": token}
	return store.filterOne(filter)
}

func (store *UserMongoDBStore) FindByTwoFactorToken(token string) (*domain.User, error) {
	filter := bson.M{"twoFactorToken": token}
	return store.filterOne(filter)
}

func (store *UserMongoDBStore) ActivateAccount(token string) *domain.User {
	filter := bson.M{"token": token}
	user, err := store.filterOne(filter)
	user.Activated = true
	user.Role = "user"
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

func (store *UserMongoDBStore) RecoverAccount(token string, newPassword string) (*domain.User, error) {
	filter := bson.M{"recoveryToken": token}
	user, err := store.filterOne(filter)
	user.RecoveryToken = ""
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	user.HashedPassword = string(hashedPassword)
	if err != nil {
		return user, err
	}
	_, err = store.users.ReplaceOne(
		context.TODO(),
		bson.M{"_id": user.Id},
		user,
	)
	if err != nil {
		return user, err
	}
	return user, err
}

func (store *UserMongoDBStore) PasswordlessLoginDemand(username string) (*domain.User, error) {
	filter := bson.M{"username": username}
	return store.filterOne(filter)
}

func (store *UserMongoDBStore) PasswordlessLogin(token string) (*domain.User, error) {
	filter := bson.M{"passwordToken": token}
	return store.filterOne(filter)
}

func (store *UserMongoDBStore) ChangePassword(username string, newPassword string, oldPassword string) error {
	filter := bson.M{"username": username}
	user, err := store.filterOne(filter)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(oldPassword))
	if err != nil {
		return err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	user.HashedPassword = string(hashedPassword)
	if err != nil {
		return err
	}
	_, err = store.users.ReplaceOne(
		context.TODO(),
		bson.M{"_id": user.Id},
		user,
	)
	if err != nil {
		return err
	}

	return nil
}

func (store *UserMongoDBStore) GenerateApiToken(username string, password string) (*domain.User, error) {
	user, err := store.filterOne(bson.M{"username": username})
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (store *UserMongoDBStore) CheckApiToken(token string) (bool, error) {
	_, err := store.filterOne(bson.M{"apiToken": token})
	if err != nil {
		return false, err
	}
	return true, nil

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

func (store *UserMongoDBStore) filterOneBlock(filter interface{}) (block *domain.Block, err error) {
	result := store.blocks.FindOne(context.TODO(), filter)
	err = result.Decode(&block)
	return
}

func (store *UserMongoDBStore) filterOneNotification(filter interface{}) (notification *domain.Notification, err error) {
	result := store.notifications.FindOne(context.TODO(), filter)
	err = result.Decode(&notification)
	return
}

func (store *UserMongoDBStore) filterBlock(filter interface{}) ([]*domain.Block, error) {
	cursor, err := store.blocks.Find(context.TODO(), filter)
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, context.TODO())

	if err != nil {
		return nil, err
	}
	return decodeBlock(cursor)
}

func (store *UserMongoDBStore) filterNotifications(filter interface{}) ([]*domain.Notification, error) {
	cursor, err := store.notifications.Find(context.TODO(), filter)
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, context.TODO())

	if err != nil {
		return nil, err
	}
	return decodeNotification(cursor)
}

func decodeNotification(cursor *mongo.Cursor) (notifications []*domain.Notification, err error) {
	for cursor.Next(context.TODO()) {
		var Notification domain.Notification
		err = cursor.Decode(&Notification)
		if err != nil {
			return
		}
		notifications = append(notifications, &Notification)
	}
	err = cursor.Err()
	return
}

func decodeBlock(cursor *mongo.Cursor) (blocks []*domain.Block, err error) {
	for cursor.Next(context.TODO()) {
		var Block domain.Block
		err = cursor.Decode(&Block)
		if err != nil {
			return
		}
		blocks = append(blocks, &Block)
	}
	err = cursor.Err()
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
