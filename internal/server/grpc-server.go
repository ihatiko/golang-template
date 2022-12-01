package server

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/ihatiko/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"net"
	grpc_container "test/internal/server/registry/components/delivery/grpc"
	"time"
)

func (s *Server) StartGrpcServer() {
	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: s.Config.Server.MaxConnectionIdle * time.Minute,
			Timeout:           s.Config.Server.TimeOut * time.Second,
			MaxConnectionAge:  s.Config.Server.MaxConnectionAge * time.Minute,
			Time:              s.Config.Server.TimeOut * time.Minute,
		}),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_opentracing.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			grpcrecovery.UnaryServerInterceptor(),
		)),
		//TODO StreamClientMetrics
	)
	listener, err := net.Listen("tcp", s.Config.Server.GrpcPort)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		err := grpcServer.Serve(listener)
		if err != nil {
			log.Fatal(err)
		}
	}()
	container := grpc_container.NewGrpcContainer(grpcServer)
	container.ServicePoints()
	grpc_prometheus.Register(grpcServer)

	s.GrpcServer = grpcServer
}
