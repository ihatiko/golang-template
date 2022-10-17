package server

import (
	"github.com/ihatiko/log"
	"os"
	"os/signal"
	"syscall"
)

func (s *Server) GracefulShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	select {
	case v := <-quit:
		log.ErrorF("signal.Notify: %v", v)
	case done := <-s.Context.Done():
		log.ErrorF("ctx.Done: %v", done)
	}

	log.Info("Start stopping http server")
	if err := s.HttpServer.Shutdown(); err != nil {
		log.Fatal(err)
	}
	log.Info("Stop http server")

	log.Info("Start stopping redis client")
	if err := s.Providers.Redis.Close(); err != nil {
		log.Fatal(err)
	}
	log.Info("Stop redis client")

	log.Info("Server exited properly")
}
