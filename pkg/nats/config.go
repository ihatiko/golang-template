package nats

import "time"

type Config struct {
	URL            string
	ClusterID      string
	ClientID       string
	ProcessTimeout time.Duration
}
