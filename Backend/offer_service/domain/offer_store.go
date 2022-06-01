package domain

type OfferStore interface {
	Get(offerId string) (*Offer, error)
	GetAll() ([]*Offer, error)
	Create(offer *Offer) error
	Update(offerId string, offer *Offer) error
	DeleteAll() error
}
