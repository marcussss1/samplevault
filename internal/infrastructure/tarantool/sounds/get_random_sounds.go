package sounds

import (
	"context"
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
)

const a = `function get_random_records()
			local space = box.space.sounds
			local all_records = space:select()
			local random_records = {}
			local random_indexes = {}
		
			for i = 1, 5 do
				table.insert(random_indexes, math.random(1, #all_records))
			end
		
			for _, index in ipairs(random_indexes) do
				table.insert(random_records, all_records[index])
			end
		
			return random_records
		end`

func (r Repository) GetRandomSounds(ctx context.Context) ([]model.Sound, error) {
	var sounds []model.Sound

	resp, err := r.conn.Call("get_random_records", nil)
	if err != nil {
		return nil, fmt.Errorf("select random sounds from Tarantool storage: %w", err)
	}

	fmt.Println(resp.Data)

	return sounds, nil
}
