package files

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/marcussss1/simplevault/internal/model"
	"mime/multipart"
	"net/url"
	"path/filepath"
	"time"
)

func (s Service) UploadSample(ctx context.Context, file multipart.File, header *multipart.FileHeader, userID string) (model.Sample, error) {
	extension := filepath.Ext(header.Filename)
	filename := uuid.NewString() + extension

	url, err := s.minioRepository.UploadSample(ctx, file, filename, header.Size)
	if err != nil {
		return model.Sample{}, fmt.Errorf("upload sample from minio repository: %w", err)
	}

	sample := newSample(url, userID, filename)

	err = s.tarantoolRepository.StoreSample(ctx, sample)
	if err != nil {
		return model.Sample{}, fmt.Errorf("upload sample from minio repository: %w", err)
	}

	return sample, nil
}

func newSample(url *url.URL, userID, filename string) model.Sample {
	url.Scheme = "https"
	url.Host = "samplevault.ru"

	return model.Sample{
		ID:                uuid.NewString(),
		AuthorID:          userID,
		AudioURL:          url.String(),
		IconURL:           "default icon url",
		FileName:          filename,
		Title:             "default title",
		Duration:          "default duration",
		MusicalInstrument: "default musical instruments",
		Genre:             "default genre",
		IsFavourite:       false,
		CreatedAt:         time.Now().String(),
	}
}
