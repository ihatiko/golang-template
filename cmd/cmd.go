package cmd

import (
	"github.com/ihatiko/config"
	"github.com/ihatiko/log"
	"github.com/opentracing/opentracing-go"
	cfg "test/config"
	"test/internal/server"
	"test/internal/server/registry/providers"
	"test/pkg/minio"
)

const (
	configPath = "./config/config.yml"
)

func Run() {
	cfg, err := config.GetConfig[cfg.Config](configPath)
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

	minioClient, err := minio.NewClient(cfg.Minio)
	if err != nil {
		log.Fatal(err)
	}
	server := server.NewServer(
		cfg,
		providers.NewProvidersContainer(
			minioClient,
			cfg.FileServiceConfig,
		),
	)
	server.Run()
}
