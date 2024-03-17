package minio

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"mime/multipart"
	"path/filepath"
)

func (r *Repository) UploadSample(ctx context.Context, file multipart.File, header *multipart.FileHeader) error {
	extension := filepath.Ext(header.Filename)
	_, err := r.conn.PutObject(ctx, "samples", "sample"+extension, file, header.Size, minio.PutObjectOptions{
		ContentType: "audio/mpeg",
	})
	if err != nil {
		return fmt.Errorf("put object from minio storage: %w", err)
	}
	r.extension = extension

	return nil
}
