package sounds

import (
	"github.com/google/uuid"
	"github.com/marcussss1/simplevault/internal/model"
	"time"
)

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
