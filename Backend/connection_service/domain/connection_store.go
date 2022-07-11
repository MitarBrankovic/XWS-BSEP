package domain

import pb "dislinkt/common/proto/connection_service"

type ConnectionStore interface {
	Get(userId string) ([]*Connection, error)
	GetAll() ([]*Connection, error)
	Create(connection *Connection) (*Connection, error)
	Delete(id string) error
	DeleteAll() error
	Update(id string) (*Connection, error)
	UpdateUser(username string, user *pb.User) (*User, error)
	RecommendFriend(username string) ([]string, error)
}
