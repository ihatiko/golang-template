package server

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"sync"
	"test/config"
	feature_components "test/internal/server/registry/components/feature-components"
	"test/internal/server/registry/providers"
)

type Server struct {
	HttpServer      *fiber.App
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
	s.GracefulShutdown()
}

func (s *Server) Registry() {
	feature_components.Registry()
}
