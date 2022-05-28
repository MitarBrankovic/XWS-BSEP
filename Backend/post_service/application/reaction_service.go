package application

import "dislinkt/post_service/domain"

type ReactionService struct {
	store domain.ReactionStore
}

func NewReactionService(store domain.ReactionStore) *ReactionService {
	return &ReactionService{
		store: store,
	}
}

func (service *ReactionService) Get(reactionId string) (*domain.Reaction, error) {
	return service.store.Get(reactionId)
}

func (service *ReactionService) GetAll() ([]*domain.Reaction, error) {
	return service.store.GetAll()
}

func (service *ReactionService) Create(reaction *domain.Reaction, postId string) error {
	return service.store.Create(reaction, postId)
}

func (service *ReactionService) Update(reactionId string, reaction *domain.Reaction) error {
	return service.store.Update(reactionId, reaction)
}
