package grpc

import (
	"context"

	"service-2/proto"
)

// Service2 .
type Service2 struct {
	proto.UnimplementedService2Server
}

// Hello .
func (s Service2) Hello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	resp := &proto.HelloResponse{
		Message: "Hi! I'm service 2",
	}

	return resp, nil
}
