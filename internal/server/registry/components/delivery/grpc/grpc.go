package grpc

import (
	"google.golang.org/grpc"
)

type grpcContainer struct {
	App *grpc.Server
}

func NewOpenApiContainer(
	app *grpc.Server,
) *grpcContainer {

	return &grpcContainer{
		App: app,
	}
}
