package domain

type ReactionStore interface {
	Get(reactionId string) (*Reaction, error)
	GetAll() ([]*Reaction, error)
	Create(reaction *Reaction) error
	Update(reactionId string, reaction *Reaction) error
	DeleteAll() error
}
