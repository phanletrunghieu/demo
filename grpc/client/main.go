package main

import (
	"context"
	"log"
	"time"

	domain "github.com/phanletrunghieu/demo/grpc/proto/domain"
	pb "github.com/phanletrunghieu/demo/grpc/proto/service"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &domain.HelloRequest{Name: "Hieu Dep Trai"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
