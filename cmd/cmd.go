package cmd

import (
	"github.com/ihatiko/config"
	"github.com/ihatiko/log"
	"github.com/opentracing/opentracing-go"
	cfg "test/config"
	"test/internal/server"
	"test/internal/server/registry/providers"
)

func Run() {
	cfg, err := config.GetConfig[cfg.Config]()
	if err != nil {
		panic(err)
	}

	cfg.Log.SetConfiguration(cfg.Server.Name)
	tracer, err := cfg.Jaeger.GetTracer(cfg.Server.Name)
	if err != nil {
		log.Fatal(err)
	}
	opentracing.SetGlobalTracer(tracer.Tracer)
	defer tracer.Closer.Close()
	log.Info("Jaeger connected")

	server := server.NewServer(
		cfg, providers.NewProvidersContainer(),
	)
	server.Run()
}
