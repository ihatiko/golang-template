package service

import "test/pkg/minio"

type FileService struct {
	minio *minio.Client
}

func NewFileService(minio *minio.Client) *FileService {
	return &FileService{minio: minio}
}

func (s FileService) SaveImage(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (s FileService) Domain1Post() error {
	return nil
}
