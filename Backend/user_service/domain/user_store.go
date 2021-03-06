package domain

type UserStore interface {
	Get(userId string) (*User, error)
	GetAll() ([]*User, error)
	GetAllBlock() ([]*Block, error)
	Create(user *User) error
	Update(userId string, user *User) error
	DeleteAll() error
	DeleteAllBlocks() error
	Find(username string) (*User, error)
	FindByEmail(email string) (*User, error)
	ActivateAccount(token string) *User
	PasswordlessLoginDemand(username string) (*User, error)
	PasswordlessLogin(token string) (*User, error)
	RecoverAccount(token string, newPassword string) (*User, error)
	ChangePassword(username string, password string, password2 string) error
	GenerateApiToken(username string, password string) (*User, error)
	CheckApiToken(token string) (bool, error)
	FindByActivationToken(token string) (*User, error)
	FindByRecoveryToken(token string) (*User, error)
	FindByPasswordlessToken(token string) (*User, error)
	FindByTwoFactorToken(token string) (*User, error)
	Block(block *Block) error
	UnBlock(block *Block) error
	CreateNotification(notification *Notification) error
	DeleteNotification(notification *Notification) error
	GetNotifications(username string) ([]*Notification, error)
	DeleteAllNotifications() error
}
