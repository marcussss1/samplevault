package samples

import (
	"context"
	"github.com/minio/minio-go/v7"
)

func (s Service) GenerateSample(ctx context.Context) (*minio.Object, error) {
	return s.minioRepository.GenerateSample(ctx)
}
