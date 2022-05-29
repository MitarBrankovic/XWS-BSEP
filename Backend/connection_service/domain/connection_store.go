package domain

type ConnectionStore interface {
	Get(userId string) ([]*Connection, error)
	Create(connection *Connection) (*Connection, error)
	Delete(id string) error
	DeleteAll() error
	Update(id string) (*Connection, error)
}
