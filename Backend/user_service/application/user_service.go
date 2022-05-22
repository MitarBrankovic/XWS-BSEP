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

func (service *UserService) Find(username string) (*domain.User, error) {
	return service.store.Find(username)
}

func (service *UserService) FindByEmail(email string) (*domain.User, error) {
	return service.store.FindByEmail(email)
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

func (service *UserService) ActivateAccount(token string) *domain.User {
	return service.store.ActivateAccount(token)
}

func (service *UserService) PasswordlessLoginDemand(username string) (*domain.User, error) {
	return service.store.PasswordlessLoginDemand(username)
}

func (service *UserService) PasswordlessLogin(token string) (*domain.User, error) {
	return service.store.PasswordlessLogin(token)
}

func (service *UserService) RecoverAccount(token string, newPassword string) (*domain.User, error) {
	return service.store.RecoverAccount(token, newPassword)
}

func (service *UserService) ChangePassword(username string, newPassword string, oldPassword string) error {
	return service.store.ChangePassword(username, newPassword, oldPassword)
}
