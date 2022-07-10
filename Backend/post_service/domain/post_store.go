package domain

import (
	pb "dislinkt/common/proto/post_service"
)

type PostStore interface {
	Get(postId string) (*Post, error)
	GetAll() ([]*Post, error)
	GetLatestPost(username string) ([]*Post, error)
	Create(post *Post) error
	Update(postId string, post *Post) error
	DeleteAll() error
	GetProfilePosts(profileId string) ([]*Post, error)
	GetConnectionPosts(profileId string) ([]*Post, error)
	UpdateUser(username string, user *pb.User) (*User, error)
	GetByUser(username string) ([]*Post, error)
	AddReaction(reaction *Reaction, postId string) error
	AddComment(comment *Comment, postId string) error
}
