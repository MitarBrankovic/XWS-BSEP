package domain

type CommentStore interface {
	Get(commentId string) (*Comment, error)
	GetAll() ([]*Comment, error)
	Create(comment *Comment) error
	Update(commentId string, comment *Comment) error
	DeleteAll() error
}
