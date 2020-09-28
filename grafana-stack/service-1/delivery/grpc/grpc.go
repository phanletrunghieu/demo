package grpc

import (
	"context"

	"service-1/client/sqlite3"
	"service-1/model"
	"service-1/proto"
)

// Service1 .
type Service1 struct {
	proto.UnimplementedService1Server
}

// Hello .
func (s Service1) Hello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	resp := &proto.HelloResponse{
		Message: "Hi! I'm service 1",
	}

	db := sqlite3.GetDB(ctx)

	bank := model.Bank{
		Name: "VCB",
	}
	db.Create(&bank)

	return resp, nil
}
