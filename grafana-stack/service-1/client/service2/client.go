package service2

import (
	"time"

	otgrpc "github.com/opentracing-contrib/go-grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	clientOpentracing "service-1/client/opentracing"
	"service-1/config"
)

var conn *grpc.ClientConn
var client Service2Client

func init() {
	cfg := config.GetConfig()
	tracer := clientOpentracing.GetTracer()

	conn, err := grpc.Dial(
		cfg.Service2Host,
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

	client = NewService2Client(conn)
}

func Close() {
	conn.Close()
}

func GetClient() Service2Client {
	return client
}
