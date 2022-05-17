package api

import (
	pb "dislinkt/common/proto/post_service"
	"dislinkt/post_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapPostToPb(post *domain.Post) *pb.Post {

	pbPost := &pb.Post{
		Id:        post.Id.Hex(),
		UserId:    post.UserId.Hex(),
		CreatedAt: timestamppb.New(post.CreatedAt),
		Content: &pb.Content{
			Text:  post.Content.Text,
			Image: post.Content.Image,
		},
	}

	return pbPost
}

func mapPbToPost(pbPost *pb.Post) *domain.Post {
	user := &domain.Post{
		Id:        getObjectId(pbPost.Id),
		UserId:    getObjectId(pbPost.UserId),
		CreatedAt: pbPost.CreatedAt.AsTime(),
		Content: domain.Content{
			Text:  pbPost.Content.Text,
			Image: pbPost.Content.Image,
		},
	}

	return user
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
