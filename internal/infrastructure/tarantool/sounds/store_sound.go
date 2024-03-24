package sounds

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
)

func (r Repository) StoreSound(ctx context.Context, sound model.Sound) error {
	data := map[string]interface{}{
		"id":                 sound.ID,
		"author_id":          sound.AuthorID,
		"audio_url":          sound.AudioURL,
		"icon_url":           sound.IconURL,
		"file_name":          sound.FileName,
		"created_at":         sound.CreatedAt,
		"title":              sound.Title,
		"musical_instrument": sound.MusicalInstrument,
		"genre":              sound.Genre,
		"mood":               sound.Mood,
		"tonality":           sound.Tonality,
		"tempo":              sound.Tempo,
		"style":              sound.Style,
	}

	packedData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error packing data: %w", err)
	}

	_, err = r.conn.Insert("sounds", packedData)
	if err != nil {
		return fmt.Errorf("insert sound from tarantool storage: %w", err)
	}

	return nil
}
