package config

import (
	"github.com/ihatiko/log"
	"test/pkg/jaeger"
	"test/pkg/postgres"
	"test/pkg/redis"
	"time"
)

type Server struct {
	Name              string
	Version           string
	Port              string
	PprofPort         string
	Mode              string
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	CtxDefaultTimeout time.Duration
	Debug             bool
}

type Config struct {
	Server   *Server
	Log      *log.Config
	Postgres *postgres.Config
	Jaeger   *jaeger.Config
	Redis    *redis.Config
}
