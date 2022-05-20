package clients

import (
	"context"
	pbUser "dislinkt/common/proto/user_service"
	"google.golang.org/grpc"
	"time"
)

func NewUserClient(address string) (pbUser.UserServiceClient, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*30)
	defer cancel()
	conn, err := grpc.DialContext(ctx, address)
	if err != nil {
		return nil, err
	}
	client := pbUser.NewUserServiceClient(conn)
	return client, nil
}
