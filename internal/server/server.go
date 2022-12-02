package server

import (
	"context"
	"file_service/config"
	feature_components "file_service/internal/server/registry/components/feature-components"
	"file_service/internal/server/registry/providers"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"sync"
)

type Server struct {
	HttpServer      *fiber.App
	GrpcServer      *grpc.Server
	Config          *config.Config
	Providers       *providers.Container
	GracefulContext *GracefulContext
}

func NewServer(config *config.Config, providers *providers.Container) *Server {
	ctx, cancelFunc := context.WithCancel(context.Background())
	return &Server{
		Config:    config,
		Providers: providers,
		GracefulContext: &GracefulContext{
			Context:    ctx,
			CancelFunc: cancelFunc,
			WgJobs:     sync.WaitGroup{},
		},
	}
}

func (s *Server) Run() {
	s.Providers.Registry()
	s.Registry()
	s.StartHttpServer()
	s.StartGrpcServer()
	s.GracefulShutdown()
}

func (s *Server) Registry() {
	feature_components.Registry()
}
