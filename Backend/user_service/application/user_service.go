package application

import "dislinkt/user_service/domain"

type UserService struct {
	store domain.UserStore
}

func NewUserService(store domain.UserStore) *UserService {
	return &UserService{
		store: store,
	}
}

func (service *UserService) Get(userId string) (*domain.User, error) {
	return service.store.Get(userId)
}

func (service *UserService) GetAll() ([]*domain.User, error) {
	return service.store.GetAll()
}

func (service *UserService) Create(user *domain.User) error {
	return service.store.Create(user)
}

func (service *UserService) Update(userId string, user *domain.User) error {
	return service.store.Update(userId, user)
}
