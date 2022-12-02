package grpc

import (
	"bytes"
	"file_service/internal/features/files"
	"file_service/internal/features/files/models"
	file_service_config "file_service/pkg/file-service-config"
	"file_service/protoc/file"
	"github.com/ihatiko/log"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
)

type Handlers struct {
	Service     files.Service
	FileService *file_service_config.Config
}

func NewApiHandler(service files.Service, fileService *file_service_config.Config) *Handlers {
	return &Handlers{Service: service, FileService: fileService}
}

func (h Handlers) UploadFile(stream file.FileService_UploadFileServer) error {
	span, ctx := opentracing.StartSpanFromContext(stream.Context(), "fileGrpcHandler.UploadFile")
	defer span.Finish()
	req, err := stream.Recv()
	if err != nil {
		return status.Errorf(codes.Unknown, "cannot receive image info")
	}
	imageData := bytes.Buffer{}
	imageSize := 0

	extension := req.GetInfo().GetExtension()
	bucket := req.GetInfo().GetBucket()
	name := req.GetInfo().GetName()
	contentType := req.GetInfo().GetContentType()
	if Validate(extension, contentType, bucket, name) != nil {
		return err
	}
	for {
		log.Info("Start receive image")
		req, err := stream.Recv()
		if err == io.EOF {
			log.Info("No more data")
			break
		}
		if err != nil {
			failedRequests.Inc()
			return status.Errorf(codes.Unknown, "cannot receive chunk data: %v", err)
		}

		chunk := req.GetChunkData()
		size := len(chunk)

		log.InfoF("received a chunk with size: %d", size)

		imageSize += size
		if h.FileService.IsNotValidSize(imageSize) {
			failedRequests.Inc()
			return status.Errorf(codes.InvalidArgument, "image is too large: %d > %d", imageSize, h.FileService.MaxSizeMb)
		}
		_, err = imageData.Write(chunk)
		if err != nil {
			failedRequests.Inc()
			return status.Errorf(codes.Internal, "cannot write chunk data: %v", err)
		}
	}
	data, err := h.Service.SaveImage(ctx, models.UploadingFile{
		ContentType: contentType,
		Bucket:      bucket,
		Extension:   extension,
		Name:        name,
		Stream:      bytes.NewReader(imageData.Bytes()),
		Size:        int64(imageSize),
	})
	if err != nil {
		failedRequests.Inc()
		return status.Errorf(codes.Internal, "cannot write chunk data: %v", err)
	}
	err = stream.SendAndClose(data)
	if err != nil {
		failedRequests.Inc()
		return status.Errorf(codes.Internal, "cannot write chunk data: %v", err)
	}
	successRequests.Inc()
	log.Info("Save image success")
	return nil
}

func Validate(contentType, extension, bucket, name string) error {
	if contentType == "" {
		return status.Errorf(codes.InvalidArgument, "cannot receive contentType data: %s", contentType)
	}
	if extension == "" {
		return status.Errorf(codes.InvalidArgument, "cannot receive extension data: %s", extension)
	}
	if bucket == "" {
		return status.Errorf(codes.InvalidArgument, "cannot receive bucket data: %s", bucket)
	}
	if name == "" {
		return status.Errorf(codes.InvalidArgument, "cannot receive name data: %s", name)
	}
	return nil
}
