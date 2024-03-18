package files

import (
	"context"
	"github.com/marcussss1/simplevault/internal/model"
	"mime/multipart"
	"net/url"

	"github.com/minio/minio-go/v7"
)

type minioRepository interface {
	DownloadSample(ctx context.Context, filename string) (*minio.Object, error)
	GenerateSample(ctx context.Context) (*minio.Object, error)
	UploadSample(ctx context.Context, file multipart.File, filename string, size int64) (*url.URL, error)
}

type tarantoolRepository interface {
	StoreSample(ctx context.Context, sample model.Sample) error
}
