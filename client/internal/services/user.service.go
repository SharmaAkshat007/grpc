package services

import (
	"context"
	"grpc/client/internal/types"
	"time"

	pb "grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserServiceClient interface {
	GetUser(id int) (*types.User, error)
}

type GrpcUserServiceClient struct {
	client pb.UserClient
}

func NewGrpcUserServiceClient() (*GrpcUserServiceClient, error) {
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	c := pb.NewUserClient(conn)
	return &GrpcUserServiceClient{client: c}, nil
}

func (c *GrpcUserServiceClient) GetUser(id int) (*types.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.client.GetUser(ctx, &pb.UserRequest{UserId: int32(id)})
	if err != nil {
		return nil, err
	}
	return &types.User{UserName: r.UserName, UserEmail: r.UserEmail}, nil
}
