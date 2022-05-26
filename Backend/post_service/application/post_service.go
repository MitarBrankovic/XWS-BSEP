package application

import (
	pb "dislinkt/common/proto/post_service"
	"dislinkt/post_service/domain"
)

type PostService struct {
	store domain.PostStore
}

func NewPostService(store domain.PostStore) *PostService {
	return &PostService{
		store: store,
	}
}

func (service *PostService) Get(postId string) (*domain.Post, error) {
	return service.store.Get(postId)
}

func (service *PostService) GetAll() ([]*domain.Post, error) {
	return service.store.GetAll()
}

func (service *PostService) Create(post *domain.Post) error {
	return service.store.Create(post)
}

func (service *PostService) Update(postId string, post *domain.Post) error {
	return service.store.Update(postId, post)
}

func (service *PostService) GetProfilePosts(profileId string) ([]*domain.Post, error) {
	return service.store.GetProfilePosts(profileId)
}

func (service *PostService) GetConnectionPosts(profileId string) ([]*domain.Post, error) {
	return service.store.GetConnectionPosts(profileId)
}

func (service *PostService) UpdateProfile(username string, user *pb.User) (*domain.User, error) {
	return service.store.UpdateUser(username, user)
}
