package minio

import (
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewClient(cfg Config) (*minio.Client, error) {
	minioClient, err := minio.New(cfg.DSN, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKey, cfg.SecretKey, ""),
		Secure: false,
	})
	if err != nil {
		return nil, fmt.Errorf("create minio instance: %w", err)
	}

	return minioClient, nil
}
