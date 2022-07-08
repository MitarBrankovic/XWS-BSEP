package clients

import (
	"context"
	pbOffer "dislinkt/common/proto/offer_service"
	"google.golang.org/grpc"
	"time"
)

type OfferClient struct {
	service pbOffer.OfferServiceClient
}

func NewOfferClient(address string) (pbOffer.OfferServiceClient, error) {
	/*tlsCredentials, err := https.LoadTLSClientCredentials()
	if err != nil {
		return nil, err
	}
	opts := []grpc.DialOption{grpc.WithTransportCredentials(tlsCredentials)}*/
	opts := []grpc.DialOption{grpc.WithInsecure()}
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*30)
	defer cancel()
	conn, err := grpc.DialContext(ctx, address, opts...)
	if err != nil {
		return nil, err
	}
	client := pbOffer.NewOfferServiceClient(conn)
	return client, nil
}
