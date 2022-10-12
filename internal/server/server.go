package server

import (
	"os"
	"os/signal"
	"syscall"
	"test/config"
	feature_components "test/internal/server/registry/components/feature-components"
	"test/internal/server/registry/providers"
)

type Server struct {
	Config    *config.Config
	Providers *providers.Container
}

func NewServer(config *config.Config, providers *providers.Container) *Server {
	return &Server{Config: config, Providers: providers}
}

func (s *Server) Run() {
	s.Providers.Registry()
	s.Registry()
	s.StartHttpServer()
	s.Interrupt()
}

func (s *Server) Registry() {
	feature_components.Registry()
}

func (s *Server) Interrupt() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
}
