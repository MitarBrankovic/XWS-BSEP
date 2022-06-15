package clients

import (
	"context"
	"dislinkt/common/https"
	pbConnection "dislinkt/common/proto/connection_service"
	"google.golang.org/grpc"
	"time"
)

func NewConnectionClient(address string) (pbConnection.ConnectionServiceClient, error) {
	tlsCredentials, err := https.LoadTLSClientCredentials()
	if err != nil {
		return nil, err
	}
	opts := []grpc.DialOption{grpc.WithTransportCredentials(tlsCredentials)}
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*30)
	defer cancel()
	conn, err := grpc.DialContext(ctx, address, opts...)
	if err != nil {
		return nil, err
	}
	client := pbConnection.NewConnectionServiceClient(conn)
	return client, nil
}
