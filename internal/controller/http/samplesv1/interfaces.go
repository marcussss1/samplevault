package samplesv1

import (
	"context"
	"mime/multipart"

	"github.com/marcussss1/simplevault/internal/model"
	"github.com/minio/minio-go/v7"
)

//go:generate mockgen -source=interfaces.go -destination=./mock/interfaces.go -package=mock

type soundsService interface {
	GetUserSounds(ctx context.Context, userID string) ([]model.Sound, error)
}

type filesService interface {
	DownloadSound(ctx context.Context, filename string) (*minio.Object, error)
	GenerateSound(ctx context.Context) (*minio.Object, error)
	UploadSound(ctx context.Context, file multipart.File, header *multipart.FileHeader, userID string) (model.Sound, error)
}
