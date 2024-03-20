package audio

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
)

func (s Service) DownloadSound(ctx context.Context, filename string) (*minio.Object, error) {
	audioFile, err := s.minioRepository.DownloadSound(ctx, filename)
	if err != nil {
		return nil, fmt.Errorf("download sound from minio repository: %w", err)
	}

	return audioFile, nil
}
