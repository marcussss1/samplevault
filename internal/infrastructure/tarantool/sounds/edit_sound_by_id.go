package sounds

import (
	"context"
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
)

//nolint:gomnd
func (r Repository) EditSoundByID(ctx context.Context, sound model.EditSound) error {
	_, err := r.conn.Replace("sounds", []interface{}{
		sound.ID,
		sound.Title,
		sound.MusicalInstrument,
		sound.Genre,
		sound.Mood,
		sound.Tonality,
		sound.Tempo,
	})
	if err != nil {
		return fmt.Errorf("edit sound by id: %w", err)
	}

	return nil
}
