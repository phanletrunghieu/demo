package main

import (
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	otgrpc "github.com/opentracing-contrib/go-grpc"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"

	clientOpentracing "service-2/client/opentracing"
	serviceGRPC "service-2/delivery/grpc"
	serviceHTTP "service-2/delivery/http"
	"service-2/proto"
)

func main() {
	defer clientOpentracing.Close()
	tracer := clientOpentracing.GetTracer()

	l, err := net.Listen("tcp", ":8001")
	if err != nil {
		panic(err)
	}

	m := cmux.New(l)
	grpcL := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httpL := m.Match(cmux.HTTP1Fast())

	errs := make(chan error)

	// http
	{
		h := serviceHTTP.New()

		go func() {
			h.Listener = httpL
			errs <- h.Start("")
		}()
	}

	// grpc
	{
		s := grpc.NewServer(
			grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
				grpc_ctxtags.StreamServerInterceptor(),
				otgrpc.OpenTracingStreamServerInterceptor(tracer),
				grpc_prometheus.StreamServerInterceptor,
				grpc_recovery.StreamServerInterceptor(),
			)),
			grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
				grpc_ctxtags.UnaryServerInterceptor(),
				otgrpc.OpenTracingServerInterceptor(tracer),
				grpc_prometheus.UnaryServerInterceptor,
				grpc_recovery.UnaryServerInterceptor(),
			)),
		)

		grpcServ := &serviceGRPC.Service2{}
		proto.RegisterService2Server(s, grpcServ)

		go func() {
			errs <- s.Serve(grpcL)
		}()
	}

	go func() {
		errs <- m.Serve()
	}()

	panic(<-errs)
}
