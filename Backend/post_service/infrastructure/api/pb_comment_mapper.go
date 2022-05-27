package api

import (
	pb "dislinkt/common/proto/post_service"
	"dislinkt/post_service/domain"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapCommentToPb(comment *domain.Comment) *pb.Comment {
	pbComment := &pb.Comment{
		Id:          comment.Id.Hex(),
		User:        mapUserToPb(&comment.User),
		DateCreated: timestamppb.New(comment.DateCreated),
	}

	return pbComment
}

func mapPbToComment(pbComment *pb.Comment) *domain.Comment {
	comment := &domain.Comment{
		Id:          getObjectId(pbComment.Id),
		User:        mapPbToUser(pbComment.User),
		DateCreated: pbComment.DateCreated.AsTime(),
	}

	return comment
}
