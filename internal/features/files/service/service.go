package service

import (
	"context"
	"github.com/ihatiko/log"
	"github.com/opentracing/opentracing-go"
	"test/internal/features/files/models"
	"test/pkg/minio"
	"test/protoc/file"
)

type FileService struct {
	minio *minio.Client
}

func NewFileService(minio *minio.Client) *FileService {
	return &FileService{minio: minio}
}

func (s FileService) SaveImage(ctx context.Context, uploadingFile models.UploadingFile) (*file.UploadFileResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "fileService.UploadFile")
	defer span.Finish()
	url, err := s.minio.Put(
		ctx,
		uploadingFile.Bucket,
		uploadingFile.Name,
		uploadingFile.ContentType,
		uploadingFile.Extension,
		uploadingFile.Stream,
		uploadingFile.Size,
	)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &file.UploadFileResponse{
		Url: url,
	}, err
}
