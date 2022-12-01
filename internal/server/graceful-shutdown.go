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
		s.HttpServer.Shutdown,
		s.Delay,
		s.WaitJobs,
	)

	log.Info("Server exit properly")
}

func (s *Server) WaitJobs() error {
	s.GracefulContext.WgJobs.Wait()
	return nil
}
func (s *Server) Delay() error {
	time.Sleep(time.Second * s.Config.Server.CtxDefaultTimeout)
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
