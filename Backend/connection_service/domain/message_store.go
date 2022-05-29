package domain

type MessageStore interface {
	Get(id string) ([]*Message, error)
	Create(message *Message) (*Message, error)
	Delete(id string) error
	DeleteAll() error
	Update(id string) (*Message, error)
}
