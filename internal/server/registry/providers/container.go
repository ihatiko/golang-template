package providers

import (
	"github.com/ihatiko/di"
	file_service_config "test/pkg/file-service-config"
	"test/pkg/minio"
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
