package minio

import (
	"context"
	"github.com/minio/minio-go/v7"
	"mime/multipart"
)

func (r Repository) UploadSample(ctx context.Context, file multipart.File, header *multipart.FileHeader) error {
	_, err := r.conn.PutObject(ctx, "samples", "sample.mp3", file, header.Size, minio.PutObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}
