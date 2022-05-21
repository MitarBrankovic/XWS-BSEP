package domain

type UserStore interface {
	Get(userId string) (*User, error)
	GetAll() ([]*User, error)
	Create(user *User) error
	Update(userId string, user *User) error
	DeleteAll() error
	Find(username string) (*User, error)
	ActivateAccount(token string) *User
	PasswordlessLoginDemand(username string) (*User, error)
	PasswordlessLogin(token string) (*User, error)
}
