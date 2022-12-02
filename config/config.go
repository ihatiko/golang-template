package config

import (
	file_service_config "file_service/pkg/file-service-config"
	"file_service/pkg/jaeger"
	"file_service/pkg/minio"
	"github.com/ihatiko/log"
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
