package api

import (
	pb "dislinkt/common/proto/post_service"
	"dislinkt/post_service/domain"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapReactionToPb(reaction *domain.Reaction) *pb.Reaction {
	pbReaction := &pb.Reaction{
		Id:        reaction.Id.Hex(),
		UserId:    reaction.UserId.Hex(),
		PostId:    reaction.PostId.Hex(),
		Type:      int64(reaction.Type), //VRV JE GRESKA
		CreatedAt: timestamppb.New(reaction.CreatedAt),
	}

	return pbReaction
}

func mapPbToReaction(pbReaction *pb.Reaction) *domain.Reaction {
	reaction := &domain.Reaction{
		Id:        getObjectId(pbReaction.Id),
		UserId:    getObjectId(pbReaction.UserId),
		PostId:    getObjectId(pbReaction.PostId),
		Type:      domain.ReactionType(pbReaction.Type),
		CreatedAt: pbReaction.CreatedAt.AsTime(),
	}

	return reaction
}
