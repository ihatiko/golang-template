package grpc

import (
	"github.com/ihatiko/di"
	"google.golang.org/grpc"
	"test/protoc/file"
)

type grpcContainer struct {
	App *grpc.Server
}

func NewGrpcContainer(
	app *grpc.Server,
) *grpcContainer {

	return &grpcContainer{
		App: app,
	}
}

func (cnt *grpcContainer) ServicePoints() {
	di.Invoke(func(fileService file.FileServiceServer) {
		file.RegisterFileServiceServer(cnt.App, fileService)
	})
}
