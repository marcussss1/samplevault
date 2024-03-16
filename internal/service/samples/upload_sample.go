package samples

import (
	"context"
	"mime/multipart"
)

func (s Service) UploadSample(ctx context.Context, file multipart.File, header *multipart.FileHeader) error {
	return s.minioRepository.UploadSample(ctx, file, header)
}
