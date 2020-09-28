package main

import (
	"context"
	"log"
	"net"

	domain "github.com/phanletrunghieu/demo/grpc/proto/domain"
	pb "github.com/phanletrunghieu/demo/grpc/proto/service"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *domain.HelloRequest) (*domain.HelloResponse, error) {
	log.Printf("Received: %v", req.GetName())
	return &domain.HelloResponse{Message: "Hello " + req.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
