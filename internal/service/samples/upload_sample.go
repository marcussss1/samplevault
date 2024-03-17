package samples

import (
	"context"
	"fmt"
	"mime/multipart"
)

func (s Service) UploadSample(ctx context.Context, file multipart.File, header *multipart.FileHeader) error {
	err := s.minioRepository.UploadSample(ctx, file, header)
	if err != nil {
		return fmt.Errorf("upload sample from minio repository: %w", err)
	}

	return nil
}
