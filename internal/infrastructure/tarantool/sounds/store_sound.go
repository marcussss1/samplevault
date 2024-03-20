package sounds

import (
	"context"
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
)

func (r Repository) StoreSound(ctx context.Context, sound model.Sound) error {
	_, err := r.conn.Insert("sounds", []interface{}{
		sound.ID,
		sound.AuthorID,
		sound.AudioURL,
		sound.IconURL,
		sound.FileName,
		sound.CreatedAt,
		sound.Title,
		sound.MusicalInstrument,
		sound.Genre,
		sound.Mood,
		sound.Tonality,
		sound.Tempo,
		sound.Style,
	})
	if err != nil {
		return fmt.Errorf("insert sound from tarantool storage: %w", err)
	}

	return nil
}
