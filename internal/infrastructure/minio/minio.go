package minio

import (
	"github.com/minio/minio-go/v7"
)

type Repository struct {
	conn      *minio.Client
	extension string
}

func NewRepository(conn *minio.Client) *Repository {
	return &Repository{
		conn: conn,
	}
}
