package server

import (
	"os"
	"os/signal"
	"syscall"
	"test/config"
	"test/internal/server/registry/providers"
)

type Server struct {
	Config    *config.Config
	Providers *providers.Container
}

func NewServer(config *config.Config) *Server {
	return &Server{Config: config}
}

func (s *Server) Run() {
	s.StartHttpServer()
	s.SystemInterrupt()
}

func (s *Server) SystemInterrupt() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
}
