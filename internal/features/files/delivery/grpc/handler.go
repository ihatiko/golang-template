package grpc

import (
	"bytes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"test/internal/features/files"
	file_service_config "test/pkg/file-service-config"
	"test/protoc/file"
)

type Handlers struct {
	Service     files.Service
	FileService *file_service_config.Config
}

func NewApiHandler(service files.Service, fileService *file_service_config.Config) *Handlers {
	return &Handlers{Service: service, FileService: fileService}
}

func (h Handlers) UploadFile(stream file.FileService_UploadFileServer) error {
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
		log.Print("waiting to receive more data")
		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("no more data")
			break
		}
		if err != nil {
			return status.Errorf(codes.Unknown, "cannot receive chunk data: %v", err)
		}

		chunk := req.GetChunkData()
		size := len(chunk)

		log.Printf("received a chunk with size: %d", size)

		imageSize += size
		if h.FileService.IsNotValidSize(imageSize) {
			return status.Errorf(codes.InvalidArgument, "image is too large: %d > %d", imageSize, h.FileService.MaxSizeMb)
		}
		_, err = imageData.Write(chunk)
		if err != nil {
			return status.Errorf(codes.Internal, "cannot write chunk data: %v", err)
		}
	}
	err = h.Service.SaveImage(stream.Context())
	if err != nil {
		return status.Errorf(codes.Internal, "cannot write chunk data: %v", err)
	}
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
