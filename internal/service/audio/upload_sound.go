package audio

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/marcussss1/simplevault/internal/model"
	"mime/multipart"
	"net/url"
	"time"
)

func (s Service) UploadSound(ctx context.Context, file multipart.File, header *multipart.FileHeader, userID string) (model.Sound, error) {
	//extension := filepath.Ext(header.Filename)
	//filename := uuid.NewString() + extension

	url, err := s.minioRepository.UploadSound(ctx, file, header.Filename, header.Size)
	if err != nil {
		return model.Sound{}, fmt.Errorf("upload sound from minio repository: %w", err)
	}

	sample := newSample(url, userID, header.Filename)

	err = s.tarantoolRepository.StoreSound(ctx, sample)
	if err != nil {
		return model.Sound{}, fmt.Errorf("store sound from tarantool repository: %w", err)
	}

	return sample, nil
}

func newSample(url *url.URL, userID, filename string) model.Sound {
	url.Scheme = "https"
	url.Host = "samplevault.ru"

	return model.Sound{
		ID:                uuid.NewString(),
		AuthorID:          userID,
		AudioURL:          url.String(),
		IconURL:           "https://img.freepik.com/free-photo/the-adorable-illustration-of-kittens-playing-in-the-forest-generative-ai_260559-483.jpg?size=338&ext=jpg&ga=GA1.1.1546980028.1710892800&semt=ais",
		FileName:          filename,
		CreatedAt:         time.Now().String(),
		Title:             "title",
		MusicalInstrument: "musical instrument",
		Genre:             "genre",
		Mood:              "mood",
		Tonality:          "tonality",
		Tempo:             "tempo",
		Style:             "style",
	}
}
