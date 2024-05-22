package audio

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/marcussss1/simplevault/internal/model"
	"mime/multipart"
	"time"
)

func (s Service) UploadSound(ctx context.Context, file multipart.File, header *multipart.FileHeader, userID string, uploadSound model.UploadSound) (model.Sound, error) {
	sample := newSample(userID, "filename", uploadSound)
	err := s.tarantoolRepository.StoreSound(ctx, sample)
	if err != nil {
		return model.Sound{}, fmt.Errorf("store sound from tarantool repository: %w", err)
	}

	return sample, nil
}

func newSample(userID, filename string, uploadSound model.UploadSound) model.Sound {
	sound := model.Sound{
		ID:                uuid.NewString(),
		AuthorID:          userID,
		AudioURL:          uploadSound.AudioURL,
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
		IsGenerated:       true,
		Likes:             0,
	}
	return sound
}
