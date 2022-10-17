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

	redis, err := cfg.Redis.NewRedisClient()
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Redis connected")

	/*	natsConn, err := cfg.Nats.NewNatsConnection()
		if err != nil {
			log.FatalF("NewNatsConnect: %+v", err)
		}
		log.InfoF(
			"Nats Connected: Status: %+v IsConnected: %v ConnectedUrl: %v ConnectedServerId: %v",
			natsConn.NatsConn().Status(),
			natsConn.NatsConn().IsConnected(),
			natsConn.NatsConn().ConnectedUrl(),
			natsConn.NatsConn().ConnectedServerId(),
		)*/

	server := server.NewServer(
		cfg, providers.NewProvidersContainer(
			redis,
			nil,
		),
	)
	server.Run()
}
