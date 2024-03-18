package files

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
)

func (s Service) DownloadSample(ctx context.Context, filename string) (*minio.Object, error) {
	sampleFile, err := s.minioRepository.DownloadSample(ctx, filename)
	if err != nil {
		return nil, fmt.Errorf("download sample from minio repository: %w", err)
	}

	return sampleFile, nil
}
