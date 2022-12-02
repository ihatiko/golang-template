package files

import (
	"context"
	"file_service/internal/features/files/models"
	"file_service/protoc/file"
)

type Service interface {
	SaveImage(ctx context.Context, file models.UploadingFile) (*file.UploadFileResponse, error)
}
