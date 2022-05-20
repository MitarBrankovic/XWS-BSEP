package clients

import (
	"context"
	pbPost "dislinkt/common/proto/post_service"
	"google.golang.org/grpc"
	"time"
)

func NewPostClient(address string) (pbPost.PostServiceClient, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*30)
	defer cancel()
	conn, err := grpc.DialContext(ctx, address)
	if err != nil {
		return nil, err
	}
	client := pbPost.NewPostServiceClient(conn)
	return client, nil
}
