package clients

import (
	"context"
	pbConnection "dislinkt/common/proto/connection_service"
	"google.golang.org/grpc"
	"time"
)

func NewConnectionClient(address string) (pbConnection.ConnectionServiceClient, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*30)
	defer cancel()
	conn, err := grpc.DialContext(ctx, address)
	if err != nil {
		return nil, err
	}
	client := pbConnection.NewConnectionServiceClient(conn)
	return client, nil
}
