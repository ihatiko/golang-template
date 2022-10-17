package providers

import (
	"github.com/go-redis/redis/v8"
	"github.com/ihatiko/di"
	"github.com/nats-io/stan.go"
)

// TODO readme description
// Add various packages and providers
type Container struct {
	Redis          *redis.Client
	NatsConnection *stan.Conn
}

func NewProvidersContainer(redis *redis.Client, natsConnection *stan.Conn) *Container {
	return &Container{
		Redis:          redis,
		NatsConnection: natsConnection,
	}
}

func (c *Container) Registry() {
	di.Provide(c.Redis)
	di.Provide(c.NatsConnection)
}
