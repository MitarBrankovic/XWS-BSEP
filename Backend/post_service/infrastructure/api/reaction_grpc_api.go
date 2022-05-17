package api

import (
	"context"
	pb "dislinkt/common/proto/post_service"
	"dislinkt/post_service/application"
)

type ReactionHandler struct {
	pb.UnimplementedPostServiceServer
	service *application.ReactionService
}

func NewReactionHandler(service *application.ReactionService) *ReactionHandler {
	return &ReactionHandler{
		service: service,
	}
}

func (handler *ReactionHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponseReaction, error) {
	reactionId := request.Id
	Reaction, err := handler.service.Get(reactionId)
	if err != nil {
		return nil, err
	}
	ReactionPb := mapReactionToPb(Reaction)
	response := &pb.GetResponseReaction{
		Reaction: ReactionPb,
	}
	return response, nil
}

func (handler *ReactionHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponseReaction, error) {
	Reactions, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponseReaction{
		Reactions: []*pb.Reaction{},
	}
	for _, Reaction := range Reactions {
		current := mapReactionToPb(Reaction)
		response.Reactions = append(response.Reactions, current)
	}
	return response, nil
}

func (handler ReactionHandler) Create(ctx context.Context, request *pb.CreateRequestReaction) (*pb.CreateResponseReaction, error) {
	reaction := mapPbToReaction(request.Reaction)
	err := handler.service.Create(reaction)
	if err != nil {
		return nil, err
	}
	return &pb.CreateResponseReaction{
		Reaction: mapReactionToPb(reaction),
	}, nil
}

func (handler ReactionHandler) Update(ctx context.Context, request *pb.UpdateRequestReaction) (*pb.UpdateResponseReaction, error) {
	reaction := mapPbToReaction(request.Reaction)
	reactionId := request.Id
	err := handler.service.Update(reactionId, reaction)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateResponseReaction{
		Reaction: mapReactionToPb(reaction),
	}, nil
}
