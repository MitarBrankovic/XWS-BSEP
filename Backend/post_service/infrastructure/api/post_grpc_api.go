package api

import (
	"context"
	pb "dislinkt/common/proto/post_service"
	pbUser "dislinkt/common/proto/user_service"
	"dislinkt/post_service/application"
)

type PostHandler struct {
	pb.UnimplementedPostServiceServer
	service         *application.PostService
	reactionService *application.ReactionService
	userClient      pbUser.UserServiceClient
	commentService  *application.CommentService
}

func NewPostHandler(service *application.PostService, reactionService *application.ReactionService, commentService *application.CommentService) *PostHandler {
	return &PostHandler{
		service:         service,
		reactionService: reactionService,
		commentService:  commentService,
	}
}

func (handler *PostHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	postId := request.Id
	Post, err := handler.service.Get(postId)
	if err != nil {
		errorLog.Error("Cannot get post: %v", err)
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
		errorLog.Error("Cannot get all posts: %v", err)
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
		errorLog.Error("Cannot create post: %v", err)
		return nil, err
	}
	successLog.Info("Post created")
	return &pb.CreateResponse{
		Post: mapPostToPb(post),
	}, nil
}

func (handler PostHandler) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	post := mapPbToPost(request.Post)
	postId := request.Id
	err := handler.service.Update(postId, post)
	if err != nil {
		errorLog.Error("Cannot update post")
		return nil, err
	}
	successLog.WithField("id", postId).Info("Post updated")
	return &pb.UpdateResponse{
		Post: mapPostToPb(post),
	}, nil
}

func (handler *PostHandler) GetProfilePosts(ctx context.Context, request *pb.GetPostRequest) (*pb.GetPostsResponse, error) {
	Posts, err := handler.service.GetProfilePosts(request.ProfileId)
	if err != nil {
		errorLog.Error("Cannot get all profile posts: %v", err)
		return nil, err
	}
	response := &pb.GetPostsResponse{
		Posts: []*pb.Post{},
	}
	for _, Post := range Posts {
		current := mapPostToPb(Post)
		response.Posts = append(response.Posts, current)
	}
	return response, nil
}

func (handler *PostHandler) GetConnectionPosts(ctx context.Context, request *pb.GetPostRequest) (*pb.GetPostsResponse, error) {
	Posts, err := handler.service.GetConnectionPosts(request.ProfileId)
	if err != nil {
		errorLog.Error("Cannot get connection posts: %v", err)
		return nil, err
	}
	response := &pb.GetPostsResponse{
		Posts: []*pb.Post{},
	}
	for _, Post := range Posts {
		current := mapPostToPb(Post)
		response.Posts = append(response.Posts, current)
	}
	return response, nil
}

func (handler *PostHandler) UpdateUser(ctx context.Context, request *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	user, err := handler.service.UpdateUser(request.User.Username, request.User)
	if err != nil {
		errorLog.Error("Cannot update user: %v", err)
		return nil, err
	}
	successLog.WithField("id", user.Id).Info("User updated")
	return &pb.UpdateUserResponse{
		User: mapUserToPb(user),
	}, nil
}

func (handler *PostHandler) GetByUser(ctx context.Context, request *pb.GetByUserRequest) (*pb.GetByUserResponse, error) {
	posts, err := handler.service.GetByUser(request.Username)
	pbPost := []*pb.Post{}
	for _, post := range posts {
		pbPost = append(pbPost, mapPostToPb(post))
	}
	if err != nil {
		errorLog.Error("Cannot get by user: %v", err)
		return nil, err
	}
	return &pb.GetByUserResponse{
		UserPosts: pbPost,
	}, nil
}

func (handler *PostHandler) CreateReaction(ctx context.Context, request *pb.CreateRequestReaction) (*pb.CreateResponseReaction, error) {
	reaction := mapPbToReaction(request.Reaction)
	err := handler.reactionService.Create(reaction, request.PostId)
	if err != nil {
		errorLog.Error("Cannot create reaction: %v", err)
		return nil, err
	}
	successLog.Info("Reaction created")
	return &pb.CreateResponseReaction{
		Reaction: mapReactionToPb(reaction),
	}, nil
}

func (handler *PostHandler) CreateComment(ctx context.Context, request *pb.CreateRequestComment) (*pb.CreateResponseComment, error) {
	comment := mapPbToComment(request.Comment)
	err := handler.commentService.Create(comment, request.PostId)
	if err != nil {
		errorLog.Error("Cannot create comment: %v", err)
		return nil, err
	}
	successLog.Info("Comment created")
	return &pb.CreateResponseComment{
		Comment: mapCommentToPb(comment),
	}, nil
}
