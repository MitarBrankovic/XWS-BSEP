package application

import "dislinkt/post_service/domain"

type ReactionService struct {
	store     domain.ReactionStore
	postStore domain.PostStore
}

func NewReactionService(store domain.ReactionStore, postStore domain.PostStore) *ReactionService {
	return &ReactionService{
		store:     store,
		postStore: postStore,
	}
}

func (service *ReactionService) Get(reactionId string) (*domain.Reaction, error) {
	return service.store.Get(reactionId)
}

func (service *ReactionService) GetAll() ([]*domain.Reaction, error) {
	return service.store.GetAll()
}

func (service *ReactionService) Create(reaction *domain.Reaction, postId string) error {
	return service.postStore.AddReaction(reaction, postId)
}

func (service *ReactionService) Update(reactionId string, reaction *domain.Reaction) error {
	return service.store.Update(reactionId, reaction)
}
