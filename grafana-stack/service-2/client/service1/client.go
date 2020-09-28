package service1

import (
	"time"

	otgrpc "github.com/opentracing-contrib/go-grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	clientOpentracing "service-2/client/opentracing"
	"service-2/config"
)

var conn *grpc.ClientConn
var client Service1Client

func init() {
	cfg := config.GetConfig()
	tracer := clientOpentracing.GetTracer()

	conn, err := grpc.Dial(
		cfg.Service1Host,
		grpc.WithInsecure(),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                5 * time.Minute,
			PermitWithoutStream: true,
		}),
		grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(tracer)),
		grpc.WithStreamInterceptor(otgrpc.OpenTracingStreamClientInterceptor(tracer)),
	)

	if err != nil {
		panic(err)
	}

	client = NewService1Client(conn)
}

func Close() {
	conn.Close()
}

func GetClient() Service1Client {
	return client
}
