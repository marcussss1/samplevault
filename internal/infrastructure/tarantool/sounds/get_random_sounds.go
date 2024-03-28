package sounds

import (
	"context"
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
)

func (r Repository) GetRandomSounds(ctx context.Context) ([]model.Sound, error) {
	query := `
		local random_sounds = {}

		for i = 0, 4 do
			local random_sound = box.space.sounds.index.primary:random(math.random(0, 1000000))
			table.insert(random_sounds, random_sound)
		end
		
		return random_sounds
	`

	resp, err := r.conn.Eval(query, []interface{}{})
	if err != nil {
		return nil, fmt.Errorf("select random sounds from tarantool storage: %w", err)
	}

	return toSounds(resp), nil
}
