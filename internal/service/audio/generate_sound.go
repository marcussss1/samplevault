package audio

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
)

func (s Service) GenerateSound(ctx context.Context) (*minio.Object, error) {
	audioFile, err := s.minioRepository.GenerateSound(ctx)
	if err != nil {
		return nil, fmt.Errorf("generate sound from minio repository: %w", err)
	}

	return audioFile, nil
}
