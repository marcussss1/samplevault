package sounds

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"mime/multipart"
	"net/url"
	"time"
)

func (r *Repository) UploadSound(ctx context.Context, file multipart.File, filename string, size int64) (*url.URL, error) {
	_, err := r.conn.PutObject(ctx, "sounds", filename, file, size, minio.PutObjectOptions{
		ContentType: "audio/mpeg",
	})
	if err != nil {
		return nil, fmt.Errorf("put object from minio storage: %w", err)
	}

	url, err := r.conn.PresignedGetObject(ctx, "sounds", filename, 7*24*time.Hour, nil)
	if err != nil {
		return nil, fmt.Errorf("presigned get object from minio storage: %w", err)
	}

	return url, nil
}
