package samplesv1

import (
	"context"
	"mime/multipart"

	"github.com/marcussss1/simplevault/internal/model"
	"github.com/minio/minio-go/v7"
)

//go:generate mockgen -source=interfaces.go -destination=./mock/interfaces.go -package=mock

type samplesService interface {
	GetAllSamples(ctx context.Context, userID string) ([]model.Sample, error)
}

type filesService interface {
	DownloadSample(ctx context.Context, filename string) (*minio.Object, error)
	GenerateSample(ctx context.Context) (*minio.Object, error)
	UploadSample(ctx context.Context, file multipart.File, header *multipart.FileHeader, userID string) (model.Sample, error)
}
