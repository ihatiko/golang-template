package server

import (
	"context"
	"github.com/ihatiko/log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type GracefulContext struct {
	CancelFunc context.CancelFunc
	Context    context.Context
	WgJobs     sync.WaitGroup
}

func (s *Server) GracefulShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	Compose(
		s.StopGrpc,
		s.StopHttp,
		s.Delay,
		s.WaitJobs,
	)

	log.Info("Server exit properly")
}

func (s *Server) StopHttp() error {
	log.Info("Stopping http server")
	return s.HttpServer.Shutdown()
}
func (s *Server) StopGrpc() error {
	log.Info("Stopping grpc server")
	s.GrpcServer.GracefulStop()
	return nil
}
func (s *Server) WaitJobs() error {
	log.Info("Waiting jobs")
	s.GracefulContext.WgJobs.Wait()
	return nil
}
func (s *Server) Delay() error {
	log.Info("Waiting timeout")
	time.Sleep(time.Second * s.Config.Server.TimeOut)
	return nil
}

func Compose(fns ...func() error) {
	for _, fn := range fns {
		err := fn()
		if err != nil {
			log.Error(err)
		}
	}
}
