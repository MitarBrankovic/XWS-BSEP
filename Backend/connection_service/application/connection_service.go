package application

import (
	pb "dislinkt/common/proto/connection_service"
	"dislinkt/connection_service/domain"
)

type ConnectionService struct {
	store domain.ConnectionStore
}

func NewConnectionService(store domain.ConnectionStore) *ConnectionService {
	return &ConnectionService{
		store: store,
	}
}

func (service *ConnectionService) Get(userId string) ([]*domain.Connection, error) {
	return service.store.Get(userId)
}

func (service *ConnectionService) GetAll() ([]*domain.Connection, error) {
	return service.store.GetAll()
}

func (service *ConnectionService) Create(connection *domain.Connection) (*domain.Connection, error) {
	return service.store.Create(connection)
}

func (service *ConnectionService) Delete(id string) error {
	return service.store.Delete(id)
}

func (service *ConnectionService) Update(id string) (*domain.Connection, error) {
	return service.store.Update(id)
}

func (service *ConnectionService) UpdateUser(username string, user *pb.User) (*domain.User, error) {
	return service.store.UpdateUser(username, user)
}
