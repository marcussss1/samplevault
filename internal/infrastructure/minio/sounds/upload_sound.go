package sounds

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"mime/multipart"
)

func (r *Repository) UploadSound(ctx context.Context, file multipart.File, filename string, size int64) error {
	_, err := r.conn.PutObject(ctx, "sounds", filename, file, size, minio.PutObjectOptions{
		ContentType: "audio/mpeg",
	})
	if err != nil {
		return fmt.Errorf("put object from minio storage: %w", err)
	}

	return nil
}
