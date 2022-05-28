package application

import "dislinkt/post_service/domain"

type ReactionService struct {
	store domain.PostStore
}

func NewReactionService(store domain.PostStore) *ReactionService {
	return &ReactionService{
		store: store,
	}
}

func (service *ReactionService) Get(reactionId string) (*domain.Reaction, error) {
	//return service.store.Get(reactionId)
	return nil, nil
}

func (service *ReactionService) GetAll() ([]*domain.Reaction, error) {
	//return service.store.GetAll()
	return nil, nil
}

func (service *ReactionService) Create(reaction *domain.Reaction, postId string) error {
	return service.store.AddReaction(reaction, postId)
}

func (service *ReactionService) Update(reactionId string, reaction *domain.Reaction) error {
	//return service.store.Update(reactionId, reaction)
	return nil
}
