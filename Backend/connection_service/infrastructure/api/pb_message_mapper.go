package api

import (
	pb "dislinkt/common/proto/connection_service"
	"dislinkt/connection_service/domain"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapMessageToPb(message *domain.Message) *pb.Message {
	return &pb.Message{
		Id:               message.Id.Hex(),
		SenderUsername:   message.SenderUsername,
		ReceiverUsername: message.ReceiverUsername,
		Date:             timestamppb.New(message.Date),
		Content:          message.Content,
	}
}

func mapPbToMessage(pbMessage *pb.Message) *domain.Message {
	return &domain.Message{
		Id:               getObjectId(pbMessage.Id),
		SenderUsername:   pbMessage.SenderUsername,
		ReceiverUsername: pbMessage.ReceiverUsername,
		Date:             pbMessage.Date.AsTime(),
		Content:          pbMessage.Content,
	}
}
