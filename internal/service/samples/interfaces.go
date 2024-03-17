package samples

import (
	"context"
	"mime/multipart"

	"github.com/marcussss1/simplevault/internal/model"
	"github.com/minio/minio-go/v7"
)

type tarantoolRepository interface {
	GetAllSamples(ctx context.Context, userID string) ([]model.Sample, error)
}

type minioRepository interface {
	DownloadSample(ctx context.Context) (*minio.Object, error)
	GenerateSample(ctx context.Context) (*minio.Object, error)
	UploadSample(ctx context.Context, file multipart.File, header *multipart.FileHeader) error
}