package sounds

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
)

func (r *Repository) DownloadSound(ctx context.Context, filename string) (*minio.Object, error) {
	sampleFile, err := r.conn.GetObject(ctx, "sounds", filename, minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("get object from minio storage: %w", err)
	}

	return sampleFile, nil
}
