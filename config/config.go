package config

import (
	"github.com/ihatiko/log"
	file_service_config "test/pkg/file-service-config"
	"test/pkg/jaeger"
	"test/pkg/minio"
	"time"
)

type Server struct {
	Name              string
	Port              string
	GrpcPort          string
	Mode              string
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	TimeOut           time.Duration
	MaxConnectionIdle time.Duration
	MaxConnectionAge  time.Duration
}

type Config struct {
	Server            *Server
	Log               *log.Config
	Jaeger            *jaeger.Config
	Minio             *minio.Config
	FileServiceConfig *file_service_config.Config
}
