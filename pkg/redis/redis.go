package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type Config struct {
	Host        string
	Password    string
	DefaultDb   string
	MinIdleConn int
	PoolSize    int
	PoolTimeout int
	DB          int
}

func (cfg *Config) NewRedisClient() (*redis.Client, error) {
	redisHost := cfg.Host

	if redisHost == "" {
		redisHost = ":6379"
	}

	client := redis.NewClient(&redis.Options{
		Addr:         redisHost,
		MinIdleConns: cfg.MinIdleConn,
		PoolSize:     cfg.PoolSize,
		PoolTimeout:  time.Duration(cfg.PoolTimeout) * time.Second,
		Password:     cfg.Password,
		DB:           cfg.DB,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return client, nil
}
