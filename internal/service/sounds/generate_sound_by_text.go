package sounds

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/marcussss1/simplevault/internal/model"
	"time"
)

func (s Service) GenerateSoundByText(ctx context.Context, text string) (model.Sound, error) {
	audioFile, err := s.mlClient.GenerateSoundByText(ctx, text)
	if err != nil {
		return model.Sound{}, fmt.Errorf("genereate sound by text from ml client: %w", err)
	}

	audioFileInfo, err := audioFile.Stat()
	if err != nil {
		return model.Sound{}, fmt.Errorf("get audio file info: %w", err)
	}

	err = s.minioRepository.UploadSound(ctx, audioFile, audioFile.Name(), audioFileInfo.Size())
	if err != nil {
		return model.Sound{}, fmt.Errorf("upload sound from minio repository: %w", err)
	}

	sample := newSample("generated_user_id", audioFile.Name())
	err = s.tarantoolRepository.StoreSound(ctx, sample)
	if err != nil {
		return model.Sound{}, fmt.Errorf("store sound from tarantool repository: %w", err)
	}

	return sample, nil
}

func newSample(userID, filename string) model.Sound {
	return model.Sound{
		ID:                uuid.NewString(),
		AuthorID:          userID,
		AudioURL:          "https://samplevault.ru/sounds/" + filename,
		IconURL:           "https://img.freepik.com/free-photo/the-adorable-illustration-of-kittens-playing-in-the-forest-generative-ai_260559-483.jpg?size=338&ext=jpg&ga=GA1.1.1546980028.1710892800&semt=ais",
		FileName:          filename,
		CreatedAt:         time.Now().String(),
		Title:             "empty",
		MusicalInstrument: "empty",
		Genre:             "empty",
		Mood:              "empty",
		Tonality:          "empty",
		Tempo:             "empty",
		Style:             "empty",
		IsGenerated:       true,
	}
}
