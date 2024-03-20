package sounds

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
)

func (r *Repository) GenerateSound(ctx context.Context) (*minio.Object, error) {
	sampleFile, err := r.conn.GetObject(ctx, "sounds", "sample.mp3", minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("get object from minio storage: %w", err)
	}

	return sampleFile, nil
}
