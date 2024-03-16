package samples

import (
	"context"
	"github.com/minio/minio-go/v7"
)

func (s Service) DownloadSample(ctx context.Context) (*minio.Object, error) {
	return s.minioRepository.DownloadSample(ctx)
}
