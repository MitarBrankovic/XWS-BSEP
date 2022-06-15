package clients

import (
	"context"
	"dislinkt/common/https"
	pbUser "dislinkt/common/proto/user_service"
	"google.golang.org/grpc"
	"time"
)

type UserClient struct {
	service pbUser.UserServiceClient
}

func NewUserClient(address string) (pbUser.UserServiceClient, error) {
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
	client := pbUser.NewUserServiceClient(conn)
	return client, nil
}
