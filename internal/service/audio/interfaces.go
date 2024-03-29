package audio

import (
	"context"
	"github.com/marcussss1/simplevault/internal/model"
	"mime/multipart"
	"net/url"

	"github.com/minio/minio-go/v7"
)

type minioRepository interface {
	DownloadSound(ctx context.Context, filename string) (*minio.Object, error)
	GenerateSound(ctx context.Context) (*minio.Object, error)
	UploadSound(ctx context.Context, file multipart.File, filename string, size int64) (*url.URL, error)
}

type tarantoolRepository interface {
	StoreSound(ctx context.Context, sound model.Sound) error
}
