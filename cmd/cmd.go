package cmd

import (
	"github.com/ihatiko/config"
	"github.com/ihatiko/log"
	"github.com/opentracing/opentracing-go"
	cfg "test/config"
	"test/internal/server"
	"test/internal/server/registry/providers"
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

	redis, err := cfg.Redis.NewRedisClient()
	if err != nil {
		log.Fatal(err)
	}
	log.Info("redis connected")

	postgres, err := cfg.Postgres.NewConnection()
	if err != nil {
		log.Fatal(err)
	}
	log.Info("postgres connected")

	natsConn, err := cfg.Nats.NewNatsConnection()
	if err != nil {
		log.FatalF("NewNatsConnect: %+v", err)
	}
	log.InfoF(
		"Nats Connected: Status: %+v",
		natsConn.NatsConn().Status(),
		natsConn.NatsConn().IsConnected(),
		natsConn.NatsConn().ConnectedUrl(),
		natsConn.NatsConn().ConnectedServerId(),
	)

	server := server.NewServer(
		cfg, providers.NewProvidersContainer(
			redis,
			&natsConn,
			postgres,
		),
	)
	server.Run()
}
