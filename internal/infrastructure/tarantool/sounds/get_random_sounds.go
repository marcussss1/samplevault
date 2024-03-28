package sounds

import (
	"context"
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
)

func (r Repository) GetRandomSounds(ctx context.Context) ([]model.Sound, error) {
	var sounds []model.Sound

	resp, err := r.conn.Eval(`
		local count = 0
		local random_sounds = {}
		
		for _, tuple in box.space.sounds.index.primary:pairs() do
			if count < 5 then
				table.insert(random_sounds, tuple)
				count = count + 1
			else
				break
			end
		end
		
		return random_sounds
	`, []interface{}{})
	if err != nil {
		return nil, fmt.Errorf("select random sounds from tarantool storage: %w", err)
	}

	for _, tuples := range resp.Tuples() {
		for _, tuple := range tuples {
			switch t := tuple.(type) {
			case []interface{}:
				sounds = append(sounds, model.Sound{
					ID:                fmt.Sprint(t[0]),
					AuthorID:          fmt.Sprint(t[1]),
					AudioURL:          fmt.Sprint(t[2]),
					IconURL:           fmt.Sprint(t[3]),
					FileName:          fmt.Sprint(t[4]),
					CreatedAt:         fmt.Sprint(t[5]),
					Title:             fmt.Sprint(t[6]),
					MusicalInstrument: fmt.Sprint(t[7]),
					Genre:             fmt.Sprint(t[8]),
					Mood:              fmt.Sprint(t[9]),
					Tonality:          fmt.Sprint(t[10]),
					Tempo:             fmt.Sprint(t[11]),
					Style:             fmt.Sprint(t[12]),
					IsGenerated:       false,
				})
			default:
				panic(1)
			}
		}
	}

	return sounds, nil
}
