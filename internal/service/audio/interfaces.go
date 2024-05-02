package audio

import (
	"context"
	"github.com/marcussss1/simplevault/internal/model"
	"github.com/minio/minio-go/v7"
	"mime/multipart"
)

type minioRepository interface {
	DownloadSound(ctx context.Context, filename string) (*minio.Object, error)
	UploadSound(ctx context.Context, file multipart.File, filename string, size int64) error
}

type tarantoolRepository interface {
	StoreSound(ctx context.Context, sound model.Sound) error
}
