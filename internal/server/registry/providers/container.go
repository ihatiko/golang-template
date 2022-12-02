package providers

import (
	file_service_config "file_service/pkg/file-service-config"
	"file_service/pkg/minio"
	"github.com/ihatiko/di"
)

type Container struct {
	Minio             *minio.Client
	FileServiceConfig *file_service_config.Config
}

func NewProvidersContainer(
	minio *minio.Client,
	fileServiceConfig *file_service_config.Config,
) *Container {
	return &Container{
		Minio:             minio,
		FileServiceConfig: fileServiceConfig,
	}
}

func (c *Container) Registry() {
	di.Provide(c.Minio)
	di.Provide(c.FileServiceConfig)
}
