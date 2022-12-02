package cmd

import (
	cfg "file_service/config"
	"file_service/internal/server"
	"file_service/internal/server/registry/providers"
	"file_service/pkg/minio"
	"github.com/ihatiko/config"
	"github.com/ihatiko/log"
	"github.com/opentracing/opentracing-go"
)

const (
	configPath = "./config/config"
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
