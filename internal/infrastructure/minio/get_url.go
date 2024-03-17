package minio

import (
	"context"
	"fmt"
	"time"
)

func (r *Repository) GetUrl() (string, error) {
	url, err := r.conn.PresignedGetObject(
		context.Background(),
		"samples",
		"sample"+r.extension,
		30*time.Minute,
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("get object from minio storage: %w", err)
	}

	return url.String(), nil
}
