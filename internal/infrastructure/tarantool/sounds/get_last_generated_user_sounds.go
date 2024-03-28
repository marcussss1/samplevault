package sounds

import (
	"context"
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
)

func (r Repository) GetLastGeneratedUserSounds(ctx context.Context, userID string) ([]model.Sound, error) {
	query := `
		local function compare(a, b)
			return a.created_at < b.created_at
		end

		local function firstFive(generated_sounds)
			local result = {}
			for i = 0, 4 do
				table.insert(result, generated_sounds[i])
			end
			return result
		end

		local generated_sounds = box.space.sounds.index.is_generated:select(true,{})
	
		local generated_user_sounds = {}
		for _, sound in ipairs(generated_sounds) do
	    	if sound.author_id == userID then
				table.insert(generated_user_sounds, sound)
	  		end
		end

		table.sort(generated_user_sounds, compare)

		return firstFive(generated_user_sounds)
	`

	resp, err := r.conn.Eval(query, []interface{}{userID})
	if err != nil {
		return nil, fmt.Errorf("select random sounds from tarantool storage: %w", err)
	}

	return toSounds(resp), nil
}
