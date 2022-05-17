package domain

type UserStore interface {
	Get(userId string) (*User, error)
	GetAll() ([]*User, error)
	Create(user *User) error
	Update(userId string, user *User) error
	DeleteAll() error
}
