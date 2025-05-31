package services

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"siakad-digi/internal/constant"
	"siakad-digi/internal/exception"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

type UploadService struct {
	MinioClient *minio.Client
}

func (s *UploadService) UploadImage(image *multipart.FileHeader) string {
	buffer, _ := image.Open()

	name := uuid.NewString() + filepath.Ext(image.Filename)
	contenType := image.Header.Get("content-type")
	c := context.Background()

	_, err := s.MinioClient.PutObject(c, constant.ASSETS_BUCKET, name, buffer, image.Size, minio.PutObjectOptions{
		ContentType: contenType,
	})

	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	minioEndpoint := os.Getenv("MINIO_ENDPOINT")
	url := fmt.Sprintf("%s/%s/%s", minioEndpoint, constant.ASSETS_BUCKET, name)

	return url
}
