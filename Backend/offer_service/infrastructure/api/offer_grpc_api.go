package api

import (
	pb "dislinkt/common/proto/offer_service"
	pbUser "dislinkt/common/proto/user_service"
	"errors"

	//pbOffer "dislinkt/common/proto/offer_service"
	"context"
	"dislinkt/offer_service/application"
)

type OfferHandler struct {
	pb.UnimplementedOfferServiceServer
	service    *application.OfferService
	userClient pbUser.UserServiceClient
}

func NewOfferHandler(service *application.OfferService, userClient pbUser.UserServiceClient) *OfferHandler {
	return &OfferHandler{
		service:    service,
		userClient: userClient,
	}
}

func (handler *OfferHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	offerId := request.Id
	Offer, err := handler.service.Get(offerId)
	if err != nil {
		return nil, err
	}
	OfferPb := mapOfferToPb(Offer)
	response := &pb.GetResponse{
		Offer: OfferPb,
	}
	return response, nil
}

func (handler *OfferHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	Offers, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Offers: []*pb.Offer{},
	}
	for _, Offer := range Offers {
		current := mapOfferToPb(Offer)
		response.Offers = append(response.Offers, current)
	}
	return response, nil
}

func (handler OfferHandler) Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	offer := mapPbToOffer(request.Offer)
	err := handler.service.Create(offer)
	if err != nil {
		return nil, err
	}
	return &pb.CreateResponse{
		Offer: mapOfferToPb(offer),
	}, nil
}

func (handler OfferHandler) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	offer := mapPbToOffer(request.Offer)
	offerId := request.Id
	err := handler.service.Update(offerId, offer)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateResponse{
		Offer: mapOfferToPb(offer),
	}, nil
}

func (handler OfferHandler) CreateMono(ctx context.Context, request *pb.CreateMonoRequest) (*pb.CreateMonoResponse, error) {
	offer := mapPbToOffer(request.Offer)
	_, err := handler.userClient.CheckApiToken(context.Background(), &pbUser.CheckApiTokenRequest{Token: request.Token})
	if err != nil {
		return nil, err
	}
	if request.Token == "" {
		return nil, errors.New("token is empty")
	}
	err = handler.service.Create(offer)
	if err != nil {
		return nil, err
	}
	return &pb.CreateMonoResponse{
		Offer: mapOfferToPb(offer),
	}, nil
}
