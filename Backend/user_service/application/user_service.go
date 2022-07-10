package application

import (
	commonDomain "dislinkt/common/domain"
	"dislinkt/user_service/domain"
)

type UserService struct {
	store        domain.UserStore
	orchestrator *UpdateUserOrchestrator
}

func NewUserService(store domain.UserStore, orchestrator *UpdateUserOrchestrator) *UserService {
	return &UserService{
		store:        store,
		orchestrator: orchestrator,
	}
}

func (service *UserService) RollbackUpdate(user *domain.User) error {
	return service.store.Update(user.Id.Hex(), user)
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

func (service *UserService) Block(block *domain.Block) error {
	return service.store.Block(block)
}

func (service *UserService) UnBlock(block *domain.Block) error {
	return service.store.UnBlock(block)
}

func (service *UserService) GetAllBlock() ([]*domain.Block, error) {
	return service.store.GetAllBlock()
}

func (service *UserService) CreateNotification(notification *domain.Notification) error {
	return service.store.CreateNotification(notification)
}

func (service *UserService) DeleteNotification(notification *domain.Notification) error {
	return service.store.DeleteNotification(notification)
}

func (service *UserService) GetNotifications(username string, notificationType domain.NotificationType) ([]*domain.Notification, error) {
	return service.store.GetNotifications(username, notificationType)
}

func (service *UserService) Update(userId string, user *domain.User) error {
	oldUser, err := service.store.Get(userId)
	if err != nil {
		return err
	}
	err = service.store.Update(userId, user)
	if err != nil {
		return err
	}
	newUser := &commonDomain.User{
		Id:        user.Id,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Private:   user.Private,
	}
	err = service.orchestrator.Start(newUser, oldUser.Username, oldUser.FirstName, oldUser.LastName, oldUser.Private)
	if err != nil {
		return err
	}
	return nil
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

func (service *UserService) GenerateApiToken(username string, password string) (*domain.User, error) {
	return service.store.GenerateApiToken(username, password)
}

func (service *UserService) CheckApiToken(token string) (bool, error) {
	return service.store.CheckApiToken(token)
}

func (service *UserService) FindByActivationToken(token string) (*domain.User, error) {
	return service.store.FindByActivationToken(token)
}

func (service *UserService) FindByRecoveryToken(token string) (*domain.User, error) {
	return service.store.FindByRecoveryToken(token)
}

func (service *UserService) FindByPasswordlessToken(token string) (*domain.User, error) {
	return service.store.FindByPasswordlessToken(token)
}

func (service *UserService) FindByTwoFactorToken(token string) (*domain.User, error) {
	return service.store.FindByTwoFactorToken(token)
}
