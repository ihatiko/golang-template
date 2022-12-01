package grpc

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	successRequests = promauto.NewCounter(prometheus.CounterOpts{
		Name: "grpc_uploaded_file_count_success",
	})
	failedRequests = promauto.NewCounter(prometheus.CounterOpts{
		Name: "grpc_uploaded_file_count_failed",
	})
)
