package api

import (
	pb "dislinkt/common/proto/post_service"
	"dislinkt/post_service/domain"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapCommentToPb(comment *domain.Comment) *pb.Comment {
	pbComment := &pb.Comment{
		Id:          comment.Id.Hex(),
		UserId:      comment.UserId.Hex(),
		PostId:      comment.PostId.Hex(),
		DateCreated: timestamppb.New(comment.DateCreated),
	}

	return pbComment
}

func mapPbToComment(pbComment *pb.Comment) *domain.Comment {
	comment := &domain.Comment{
		Id:          getObjectId(pbComment.Id),
		UserId:      getObjectId(pbComment.UserId),
		PostId:      getObjectId(pbComment.PostId),
		DateCreated: pbComment.DateCreated.AsTime(),
	}

	return comment
}
