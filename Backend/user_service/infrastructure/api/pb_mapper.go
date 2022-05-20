package api

import (
	pb "dislinkt/common/proto/user_service"
	"dislinkt/user_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapUserToPb(user *domain.User) *pb.User {
	pbUser := &pb.User{
		Id:          user.Id.Hex(),
		Username:    user.Username,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		DateOfBirth: timestamppb.New(user.DateOfBirth),
		Email:       user.Email,
	}

	return pbUser
}

func mapPbToUser(pbUser *pb.User) *domain.User {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(pbUser.Password), bcrypt.DefaultCost)
	user := &domain.User{
		Id:             getObjectId(pbUser.Id),
		Username:       pbUser.Username,
		HashedPassword: string(hashedPassword),
		FirstName:      pbUser.FirstName,
		LastName:       pbUser.LastName,
		FullName:       pbUser.FirstName + " " + pbUser.LastName,
		DateOfBirth:    pbUser.DateOfBirth.AsTime(),
		Email:          pbUser.Email,
	}

	return user
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
