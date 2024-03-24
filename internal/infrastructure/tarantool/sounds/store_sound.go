package sounds

import (
	"context"
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
)

func (r Repository) StoreSound(ctx context.Context, sound model.Sound) error {
	_, err := r.conn.Eval(`
	  function escapeString(str)
	      return str:gsub('[%-%.%+%[%]%(%)%*%?%^%$%%]', '%%%1')
	  end
	
	  function insertSound(sound)
	      local escapedData = {}
	      for key, value in pairs(sound) do
	          if type(value) == "string" then
	              escapedData[key] = escapeString(value)
	          else
	              escapedData[key] = value
	          end
	      end
	
	      box.space.sounds:insert(escapedData)
	  end
	`, nil)
	if err != nil {
		return fmt.Errorf("error defining Lua functions: %w", err)
	}

	_, err = r.conn.Call("insertSound", map[string]interface{}{
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
	})

	if err != nil {
		return fmt.Errorf("insert sound from tarantool storage: %w", err)
	}

	return nil
	//_, err := r.conn.Insert("sounds", []interface{}{
	//	sound.ID,
	//	sound.AuthorID,
	//	sound.AudioURL,
	//	sound.IconURL,
	//	sound.FileName,
	//	sound.CreatedAt,
	//	sound.Title,
	//	sound.MusicalInstrument,
	//	sound.Genre,
	//	sound.Mood,
	//	sound.Tonality,
	//	sound.Tempo,
	//	sound.Style,
	//})
	//if err != nil {
	//	return fmt.Errorf("insert sound from tarantool storage: %w", err)
	//}

	//return nil
}
