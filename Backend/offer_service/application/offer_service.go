package application

import "dislinkt/offer_service/domain"

type OfferService struct {
	store domain.OfferStore
}

func NewOfferService(store domain.OfferStore) *OfferService {
	return &OfferService{
		store: store,
	}
}

func (service *OfferService) Get(offerId string) (*domain.Offer, error) {
	return service.store.Get(offerId)
}

func (service *OfferService) GetAll() ([]*domain.Offer, error) {
	return service.store.GetAll()
}

func (service *OfferService) Create(offer *domain.Offer) error {
	return service.store.Create(offer)
}

func (service *OfferService) Update(offerId string, offer *domain.Offer) error {
	return service.store.Update(offerId, offer)
}
