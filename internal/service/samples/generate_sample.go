package samples

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
)

func (s Service) GenerateSample(ctx context.Context) (*minio.Object, error) {
	sampleFile, err := s.minioRepository.GenerateSample(ctx)
	if err != nil {
		return nil, fmt.Errorf("generate sample from minio repository: %w", err)
	}

	return sampleFile, nil
}
