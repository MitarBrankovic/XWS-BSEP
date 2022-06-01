package api

import (
	pb "dislinkt/common/proto/offer_service"
	"dislinkt/offer_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapOfferToPb(offer *domain.Offer) *pb.Offer {

	pbOffer := &pb.Offer{
		Id:          offer.Id.Hex(),
		Username:    offer.Username,
		Company:     offer.Company,
		Description: offer.Description,
		Position:    offer.Position,
		Criteria:    offer.Criteria,
		CreatedAt:   timestamppb.New(offer.CreatedAt),
	}

	return pbOffer
}

func mapPbToOffer(pbOffer *pb.Offer) *domain.Offer {
	offer := &domain.Offer{
		Id:          getObjectId(pbOffer.Id),
		Username:    pbOffer.Username,
		Company:     pbOffer.Company,
		Description: pbOffer.Description,
		Position:    pbOffer.Position,
		Criteria:    pbOffer.Criteria,
		CreatedAt:   pbOffer.CreatedAt.AsTime(),
	}

	return offer
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
