package audio

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/marcussss1/simplevault/internal/model"
	"mime/multipart"
	"path/filepath"
	"time"
)

func (s Service) UploadSound(ctx context.Context, file multipart.File, header *multipart.FileHeader, userID string, uploadSound model.UploadSound) (model.Sound, error) {
	extension := filepath.Ext(header.Filename)
	filename := uuid.NewString() + extension

	fmt.Println()
	fmt.Println(uploadSound)
	fmt.Println()

	err := s.minioRepository.UploadSound(ctx, file, filename, header.Size)
	if err != nil {
		return model.Sound{}, fmt.Errorf("upload sound from minio repository: %w", err)
	}

	sample := newSample(userID, filename, uploadSound)
	err = s.tarantoolRepository.StoreSound(ctx, sample)
	if err != nil {
		return model.Sound{}, fmt.Errorf("store sound from tarantool repository: %w", err)
	}

	return sample, nil
}

func newSample(userID, filename string, uploadSound model.UploadSound) model.Sound {
	sound := model.Sound{
		ID:                uuid.NewString(),
		AuthorID:          userID,
		AudioURL:          "https://samplevault.ru/sounds/" + filename,
		IconURL:           "https://img.freepik.com/free-photo/the-adorable-illustration-of-kittens-playing-in-the-forest-generative-ai_260559-483.jpg?size=338&ext=jpg&ga=GA1.1.1546980028.1710892800&semt=ais",
		FileName:          filename,
		CreatedAt:         time.Now().String(),
		Title:             uploadSound.Title,
		MusicalInstrument: uploadSound.MusicalInstrument,
		Genre:             uploadSound.Genre,
		Mood:              uploadSound.Mood,
		Tonality:          uploadSound.Tonality,
		Tempo:             uploadSound.Tempo,
		Style:             "style",
		IsGenerated:       false,
		Likes:             0,
	}
	if uploadSound.IsGenerated == "true" {
		sound.IsGenerated = true
	}
	return sound
}
