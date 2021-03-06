package api

import (
	"context"
	"dislinkt/common/loggers"
	pb "dislinkt/common/proto/post_service"
	"dislinkt/post_service/application"
)

var errorLog = loggers.NewErrorLogger()
var successLog = loggers.NewSuccessLogger()

type CommentHandler struct {
	pb.UnimplementedPostServiceServer
	service *application.CommentService
}

func NewCommentHandler(service *application.CommentService) *CommentHandler {
	return &CommentHandler{
		service: service,
	}
}

func (handler *CommentHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponseComment, error) {
	commentId := request.Id
	Comment, err := handler.service.Get(commentId)
	if err != nil {
		errorLog.Error("Cannot get comment: %v", err)
		return nil, err
	}
	CommentPb := mapCommentToPb(Comment)
	response := &pb.GetResponseComment{
		Comment: CommentPb,
	}
	return response, nil
}

func (handler *CommentHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponseComment, error) {
	Comments, err := handler.service.GetAll()
	if err != nil {
		errorLog.Error("Cannot get all comments: %v", err)
		return nil, err
	}
	response := &pb.GetAllResponseComment{
		Comments: []*pb.Comment{},
	}
	for _, Comment := range Comments {
		current := mapCommentToPb(Comment)
		response.Comments = append(response.Comments, current)
	}
	return response, nil
}

func (handler CommentHandler) Create(ctx context.Context, request *pb.CreateRequestComment) (*pb.CreateResponseComment, error) {
	comment := mapPbToComment(request.Comment)
	err := handler.service.Create(comment, request.PostId)
	if err != nil {
		errorLog.Error("Cannot create comment: %v", err)
		return nil, err
	}
	successLog.Info("Comment created")
	return &pb.CreateResponseComment{
		Comment: mapCommentToPb(comment),
	}, nil
}

func (handler CommentHandler) Update(ctx context.Context, request *pb.UpdateRequestComment) (*pb.UpdateResponseComment, error) {
	comment := mapPbToComment(request.Comment)
	commentId := request.Id
	err := handler.service.Update(commentId, comment)
	if err != nil {
		errorLog.Error("Cannot update comment")
		return nil, err
	}
	successLog.WithField("id", commentId).Info("Comment updated")
	return &pb.UpdateResponseComment{
		Comment: mapCommentToPb(comment),
	}, nil
}
