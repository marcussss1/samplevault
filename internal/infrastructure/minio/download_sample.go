package minio

import (
	"context"

	"github.com/minio/minio-go/v7"
)

func (r *Repository) DownloadSample(ctx context.Context) (*minio.Object, error) {
	sampleFile, err := r.conn.GetObject(ctx, "samples", "sample"+r.extension, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}

	return sampleFile, nil
}
