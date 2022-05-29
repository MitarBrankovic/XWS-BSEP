package application

import "dislinkt/connection_service/domain"

type MessageService struct {
	store domain.MessageStore
}

func NewMessageService(store domain.MessageStore) *MessageService {
	return &MessageService{
		store: store,
	}
}

func (service *MessageService) Get(id string) ([]*domain.Message, error) {
	return service.store.Get(id)
}

func (service *MessageService) Create(message *domain.Message) (*domain.Message, error) {
	return service.store.Create(message)
}

func (service *MessageService) Delete(id string) error {
	return service.store.Delete(id)
}

func (service *MessageService) Update(id string) (*domain.Message, error) {
	return nil, nil
}
