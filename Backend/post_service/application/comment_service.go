package application

import "dislinkt/post_service/domain"

type CommentService struct {
	store domain.CommentStore
}

func NewCommentService(store domain.CommentStore) *CommentService {
	return &CommentService{
		store: store,
	}
}

func (service *CommentService) Get(commentId string) (*domain.Comment, error) {
	return service.store.Get(commentId)
}

func (service *CommentService) GetAll() ([]*domain.Comment, error) {
	return service.store.GetAll()
}

func (service *CommentService) Create(comment *domain.Comment) error {
	return service.store.Create(comment)
}

func (service *CommentService) Update(commentId string, comment *domain.Comment) error {
	return service.store.Update(commentId, comment)
}
