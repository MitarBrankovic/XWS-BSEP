package clients

import (
	"context"
	pbOffer "dislinkt/common/proto/offer_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type OfferClient struct {
	service pbOffer.OfferServiceClient
}

func NewOfferClient(address string) (pbOffer.OfferServiceClient, error) {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*30)
	defer cancel()
	conn, err := grpc.DialContext(ctx, address, opts...)
	if err != nil {
		return nil, err
	}
	client := pbOffer.NewOfferServiceClient(conn)
	return client, nil
}
