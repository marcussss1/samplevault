package minio

import (
	"context"
	"mime/multipart"
	"path/filepath"

	"github.com/minio/minio-go/v7"
)

func (r *Repository) UploadSample(ctx context.Context, file multipart.File, header *multipart.FileHeader) error {
	extension := filepath.Ext(header.Filename)
	_, err := r.conn.PutObject(ctx, "samples", "sample"+extension, file, header.Size, minio.PutObjectOptions{})
	if err != nil {
		return err
	}
	r.extension = extension

	return nil
}
