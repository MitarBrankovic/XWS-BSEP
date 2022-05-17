package domain

type PostStore interface {
	Get(postId string) (*Post, error)
	GetAll() ([]*Post, error)
	Create(post *Post) error
	Update(postId string, post *Post) error
	DeleteAll() error
}
