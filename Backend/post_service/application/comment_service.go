package application

import "dislinkt/post_service/domain"

type CommentService struct {
	store domain.PostStore
}

func NewCommentService(store domain.PostStore) *CommentService {
	return &CommentService{
		store: store,
	}
}

func (service *CommentService) Get(commentId string) (*domain.Comment, error) {
	return nil, nil
}

func (service *CommentService) GetAll() ([]*domain.Comment, error) {
	return nil, nil
}

func (service *CommentService) Create(comment *domain.Comment, postId string) error {
	return service.store.AddComment(comment, postId)
}

func (service *CommentService) Update(commentId string, comment *domain.Comment) error {
	return nil
}
