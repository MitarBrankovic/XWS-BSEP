package api

import (
	pb "dislinkt/common/proto/connection_service"
	"dislinkt/connection_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapConnectionToPb(connection *domain.Connection) *pb.Connection {
	return &pb.Connection{
		Id:              connection.Id.Hex(),
		IssuerUsername:  connection.IssuerUsername,
		SubjectUsername: connection.SubjectUsername,
		Date:            timestamppb.New(connection.Date),
		IsApproved:      connection.IsApproved,
	}
}

func mapPbToConnection(pbConnection *pb.Connection) *domain.Connection {
	return &domain.Connection{
		Id:              getObjectId(pbConnection.Id),
		IssuerUsername:  pbConnection.IssuerUsername,
		SubjectUsername: pbConnection.SubjectUsername,
		Date:            pbConnection.Date.AsTime(),
		IsApproved:      pbConnection.IsApproved,
	}
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
