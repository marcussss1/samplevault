package minio

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
)

func (r *Repository) DownloadSample(ctx context.Context) (*minio.Object, error) {
	sampleFile, err := r.conn.GetObject(ctx, "samples", "sample"+r.extension, minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("get object from minio storage: %w", err)
	}

	return sampleFile, nil
}
