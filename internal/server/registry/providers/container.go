package providers

import (
	"github.com/go-redis/redis/v8"
	"github.com/ihatiko/di"
)

// TODO readme description
// Add various packages and providers
type Container struct {
	Redis *redis.Client
}

func NewProvidersContainer(redis *redis.Client) *Container {
	return &Container{
		Redis: redis,
	}
}

func (c *Container) Registry() {
	di.Provide(c.Redis)
}
