package api

import (
	pb "dislinkt/common/proto/connection_service"
	"dislinkt/connection_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapConnectionToPb(connection *domain.Connection) *pb.Connection {
	return &pb.Connection{
		Id:          connection.Id.Hex(),
		IssuerUser:  mapUserToPbUser(&connection.IssuerUser),
		SubjectUser: mapUserToPbUser(&connection.SubjectUser),
		Date:        timestamppb.New(connection.Date),
		IsApproved:  connection.IsApproved,
	}
}

func mapPbToConnection(pbConnection *pb.Connection) *domain.Connection {
	return &domain.Connection{
		Id:          getObjectId(pbConnection.Id),
		IssuerUser:  mapPbUserToUser(pbConnection.IssuerUser),
		SubjectUser: mapPbUserToUser(pbConnection.SubjectUser),
		Date:        pbConnection.Date.AsTime(),
		IsApproved:  pbConnection.IsApproved,
	}
}

func mapPbUserToUser(pbUser *pb.User) domain.User {
	return domain.User{
		Username: pbUser.Username,
		Private:  pbUser.Private,
	}
}

func mapUserToPbUser(user *domain.User) *pb.User {
	return &pb.User{
		Username: user.Username,
		Private:  user.Private,
	}
}

func mapSagaUserToPb(user *domain.User) *pb.User {
	pbUser := &pb.User{
		Username: user.Username,
		Private:  user.Private,
	}
	return pbUser
}

func mapPbToSagaUser(pbUser *pb.User) domain.User {
	user := domain.User{
		Username: pbUser.Username,
		Private:  pbUser.Private,
	}
	return user
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
