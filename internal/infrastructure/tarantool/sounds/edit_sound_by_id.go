package sounds

import (
	"context"
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
)

//nolint:gomnd
func (r Repository) EditSoundByID(ctx context.Context, sound model.EditSound) error {
	query := `
		local arg = {...} 
		local id = arg[1]
		local title = arg[2]
		local musicalInstrument = arg[3]
		local genre = arg[4]
		local mood = arg[5]
		local tonality = arg[6]
		local tempo = arg[7]
		
		return box.space.sounds:update(id, {
    		{'=', 'title', title},
    		{'=', 'musical_instrument', musicalInstrument},
    		{'=', 'genre', genre},
    		{'=', 'mood', mood},
    		{'=', 'tonality', tonality},
    		{'=', 'tempo', tempo},
		})
	`

	_, err := r.conn.Eval(query, []interface{}{
		sound.ID,
		sound.Title,
		sound.MusicalInstrument,
		sound.Genre,
		sound.Mood,
		sound.Tonality,
		sound.Tempo,
	})
	if err != nil {
		return fmt.Errorf("replace sound from tarantool storage: %w", err)
	}

	return nil
}
