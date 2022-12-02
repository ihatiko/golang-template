package feature_components

import (
	"file_service/internal/features/files"
	fileHandlers "file_service/internal/features/files/delivery/grpc"
	fileService "file_service/internal/features/files/service"
	pbFile "file_service/protoc/file"
	"github.com/ihatiko/di"
)

func Registry() {
	SetRepository()
	SetService()
	SetDelivery()
}

func SetDelivery() {
	di.ProvideInterface[pbFile.FileServiceServer](fileHandlers.NewApiHandler)
}

func SetRepository() {
}

func SetService() {
	di.ProvideInterface[files.Service](fileService.NewFileService)
}
