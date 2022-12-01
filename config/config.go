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
	Server      *Server
	Log         *log.Config
	Jaeger      *jaeger.Config
	Minio       *minio.Config
	FileService *file_service_config.Config
}
