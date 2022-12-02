package grpc

import (
	"file_service/protoc/file"
	"github.com/ihatiko/di"
	"google.golang.org/grpc"
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
