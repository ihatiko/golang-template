package cmd

import (
	"github.com/ihatiko/log"
	"github.com/opentracing/opentracing-go"
	"test/config"
	"test/internal/server"
	"test/internal/server/registry/providers"
)

func Run() {
	cfg, err := config.GetConfig()
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

	redis, err := cfg.Redis.NewRedisClient()
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Redis connected")
	server := server.NewServer(
		cfg, providers.NewProvidersContainer(
			redis,
		),
	)
	server.Run()
}
