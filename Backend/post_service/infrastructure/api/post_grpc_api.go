package api

import (
	"context"
	pb "dislinkt/common/proto/post_service"
	pbUser "dislinkt/common/proto/user_service"
	"dislinkt/post_service/application"
)

type PostHandler struct {
	pb.UnimplementedPostServiceServer
	service    *application.PostService
	userClient pbUser.UserServiceClient
}

func NewPostHandler(service *application.PostService, userClient pbUser.UserServiceClient) *PostHandler {
	return &PostHandler{
		service:    service,
		userClient: userClient,
	}
}

func (handler *PostHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	postId := request.Id
	Post, err := handler.service.Get(postId)
	if err != nil {
		return nil, err
	}
	PostPb := mapPostToPb(Post)
	response := &pb.GetResponse{
		Post: PostPb,
	}
	return response, nil
}

func (handler *PostHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	Posts, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Posts: []*pb.Post{},
	}
	for _, Post := range Posts {
		current := mapPostToPb(Post)
		response.Posts = append(response.Posts, current)
	}
	return response, nil
}

func (handler PostHandler) Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	post := mapPbToPost(request.Post)
	err := handler.service.Create(post)
	if err != nil {
		return nil, err
	}
	return &pb.CreateResponse{
		Post: mapPostToPb(post),
	}, nil
}

func (handler PostHandler) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	post := mapPbToPost(request.Post)
	postId := request.Id
	err := handler.service.Update(postId, post)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateResponse{
		Post: mapPostToPb(post),
	}, nil
}
