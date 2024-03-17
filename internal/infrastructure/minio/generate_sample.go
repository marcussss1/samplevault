package minio

import (
	"context"

	"github.com/minio/minio-go/v7"
)

func (r *Repository) GenerateSample(ctx context.Context) (*minio.Object, error) {
	objectReader, err := r.conn.GetObject(ctx, "samples", "sample"+r.extension, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	return objectReader, nil
}
