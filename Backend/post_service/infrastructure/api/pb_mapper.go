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
		User:      mapUserToPb(&post.User),
		CreatedAt: timestamppb.New(post.CreatedAt),
		Content: &pb.Content{
			Text:  post.Content.Text,
			Image: post.Content.Image,
		},
		Comments:  mapCommentListToPb(post.Comments),
		Reactions: mapReactionListToPb(post.Reactions),
	}

	return pbPost
}

func mapPbToPost(pbPost *pb.Post) *domain.Post {
	post := &domain.Post{
		Id:        getObjectId(pbPost.Id),
		User:      mapPbToUser(pbPost.User),
		CreatedAt: pbPost.CreatedAt.AsTime(),
		Content: domain.Content{
			Text:  pbPost.Content.Text,
			Image: pbPost.Content.Image,
		},
		Comments:  mapPbToCommentList(pbPost.Comments),
		Reactions: mapPbToReactionList(pbPost.Reactions),
	}

	return post
}

func mapUserToPb(user *domain.User) *pb.User {
	pbUser := &pb.User{
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
	return pbUser
}

func mapPbToUser(pbUser *pb.User) domain.User {
	user := domain.User{
		Username:  pbUser.Username,
		FirstName: pbUser.FirstName,
		LastName:  pbUser.LastName,
	}
	return user
}

func mapCommentListToPb(comments []domain.Comment) []*pb.Comment {
	pbComments := make([]*pb.Comment, len(comments))
	for i, comment := range comments {
		pbComments[i] = mapCommentToPb(&comment)
	}
	return pbComments
}

func mapPbToCommentList(pbComments []*pb.Comment) []domain.Comment {
	comments := make([]domain.Comment, len(pbComments))
	for i, pbComment := range pbComments {
		comments[i] = *mapPbToComment(pbComment)
	}
	return comments
}

func mapReactionListToPb(reactions []domain.Reaction) []*pb.Reaction {
	pbReactions := make([]*pb.Reaction, len(reactions))
	for i, reaction := range reactions {
		pbReactions[i] = mapReactionToPb(&reaction)
	}
	return pbReactions
}

func mapPbToReactionList(pbReactions []*pb.Reaction) []domain.Reaction {
	reactions := make([]domain.Reaction, len(pbReactions))
	for i, pbReaction := range pbReactions {
		reactions[i] = *mapPbToReaction(pbReaction)
	}
	return reactions
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
