package minio

import (
	"fmt"

	"github.com/minio/minio-go/v7"
)

type Repository struct {
	conn      *minio.Client
	extension string
}

func NewRepository(conn *minio.Client) (*Repository, error) {
	if conn == nil {
		return nil, fmt.Errorf("conn is nil")
	}

	return &Repository{
		conn: conn,
	}, nil
}
