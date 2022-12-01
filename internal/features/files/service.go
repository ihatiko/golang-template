package files

import (
	"context"
	"test/internal/features/files/models"
	"test/protoc/file"
)

type Service interface {
	SaveImage(ctx context.Context, file models.UploadingFile) (*file.UploadFileResponse, error)
}
