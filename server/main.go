package main

import (
	"context"
	"log"
	"net"

	pb "grpc/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServer
}

func (s *server) GetUser(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	log.Printf("Received: %v", in.GetUserId())
	return &pb.UserResponse{UserName: "Akshat", UserEmail: "abc@gmail.com"}, nil
}

func main() {

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterUserServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
