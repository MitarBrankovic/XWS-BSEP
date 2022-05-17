package api

import (
	"context"
	pb "dislinkt/common/proto/user_service"
	"dislinkt/user_service/application"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	service *application.UserService
}

func NewUserHandler(service *application.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (handler *UserHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	userId := request.Id
	User, err := handler.service.Get(userId)
	if err != nil {
		return nil, err
	}
	UserPb := mapUserToPb(User)
	response := &pb.GetResponse{
		User: UserPb,
	}
	return response, nil
}

func (handler *UserHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	Users, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Users: []*pb.User{},
	}
	for _, User := range Users {
		current := mapUserToPb(User)
		response.Users = append(response.Users, current)
	}
	return response, nil
}

func (handler UserHandler) Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	user := mapPbToUser(request.User)
	err := handler.service.Create(user)
	if err != nil {
		return nil, err
	}
	return &pb.CreateResponse{
		User: mapUserToPb(user),
	}, nil
}

func (handler UserHandler) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	user := mapPbToUser(request.User)
	userId := request.Id
	err := handler.service.Update(userId, user)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateResponse{
		User: mapUserToPb(user),
	}, nil
}
