package clients

import (
	"context"
	pbPost "dislinkt/common/proto/post_service"
	"google.golang.org/grpc"
	"time"
)

func NewPostClient(address string) (pbPost.PostServiceClient, error) {
	/*tlsCredentials, err := https.LoadTLSClientCredentials()
	if err != nil {
		return nil, err
	}
	opts := []grpc.DialOption{grpc.WithTransportCredentials(tlsCredentials)}*/
	opts := []grpc.DialOption{grpc.WithInsecure()}
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*30)
	defer cancel()
	/*dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}*/
	conn, err := grpc.DialContext(ctx, address, opts...)
	if err != nil {
		return nil, err
	}
	client := pbPost.NewPostServiceClient(conn)
	return client, nil
}
