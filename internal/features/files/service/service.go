package service

import (
	"context"
	"file_service/internal/features/files/models"
	"file_service/pkg/minio"
	"file_service/protoc/file"
	"github.com/ihatiko/log"
	"github.com/opentracing/opentracing-go"
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
