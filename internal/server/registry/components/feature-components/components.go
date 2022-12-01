package feature_components

import (
	"github.com/ihatiko/di"
	fileHandlers "test/internal/features/files/delivery/grpc"
	"test/protoc/file"
)

func Registry() {
	SetRepository()
	SetService()
	SetDelivery()
}

func SetDelivery() {
	di.ProvideInterface[file.FileServiceServer](fileHandlers.NewApiHandler)
}

func SetRepository() {
}

func SetService() {

}
