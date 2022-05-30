package api

import (
	"context"
	pb "dislinkt/common/proto/connection_service"
	pbUser "dislinkt/common/proto/user_service"
	"dislinkt/connection_service/application"
)

type ConnectionHandler struct {
	pb.UnimplementedConnectionServiceServer
	service        *application.ConnectionService
	messageService *application.MessageService
	userClient     pbUser.UserServiceClient
}

func NewConnectionHandler(service *application.ConnectionService, messageService *application.MessageService, userClient pbUser.UserServiceClient) *ConnectionHandler {
	return &ConnectionHandler{
		service:        service,
		messageService: messageService,
		userClient:     userClient,
	}
}

func (handler *ConnectionHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	Connections, err := handler.service.Get(request.UserId)
	if err != nil {
		return nil, err
	}
	response := &pb.GetResponse{
		Connections: []*pb.Connection{},
	}
	for _, Connection := range Connections {
		current := mapConnectionToPb(Connection)
		response.Connections = append(response.Connections, current)
	}
	return response, nil
}

func (handler *ConnectionHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	Connections, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Connections: []*pb.Connection{},
	}
	for _, Connection := range Connections {
		current := mapConnectionToPb(Connection)
		response.Connections = append(response.Connections, current)
	}
	return response, nil
}

func (handler *ConnectionHandler) Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	user, err := handler.userClient.FindByUsername(context.Background(), &pbUser.FindByUsernameRequest{Username: request.Connection.SubjectUsername})
	connection := mapPbToConnection(request.Connection)
	if user.User.Private {
		connection.IsApproved = false
	} else {
		connection.IsApproved = true
	}
	newConnection, err := handler.service.Create(connection)
	if err != nil {
		return nil, err
	}
	return &pb.CreateResponse{
		Connection: mapConnectionToPb(newConnection),
	}, nil
}

func (handler *ConnectionHandler) Delete(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	err := handler.service.Delete(request.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteResponse{}, nil
}

func (handler *ConnectionHandler) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	connection, err := handler.service.Update(request.Id)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateResponse{
		Connection: mapConnectionToPb(connection),
	}, nil
}

func (handler *ConnectionHandler) CreateMessage(ctx context.Context, request *pb.CreateMessageRequest) (*pb.CreateMessageResponse, error) {
	message := mapPbToMessage(request.Message)
	newMessage, err := handler.messageService.Create(message)
	if err != nil {
		return nil, err
	}
	return &pb.CreateMessageResponse{
		Message: mapMessageToPb(newMessage),
	}, nil
}
