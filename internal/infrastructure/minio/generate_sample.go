package minio

import (
	"context"
	"github.com/minio/minio-go/v7"
)

func (r Repository) GenerateSample(ctx context.Context) (*minio.Object, error) {
	objectReader, err := r.conn.GetObject(ctx, "samples", "sample.mp3", minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	return objectReader, nil
}
