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
	DownloadSample(ctx context.Context) (*minio.Object, error)
	GenerateSample(ctx context.Context) (*minio.Object, error)
	UploadSample(ctx context.Context, file multipart.File, header *multipart.FileHeader) error
	GetUrl() (string, error)
}
