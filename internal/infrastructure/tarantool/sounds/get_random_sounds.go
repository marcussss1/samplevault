package sounds

import (
	"context"
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
)

//nolint:gomnd
func (r Repository) GetRandomSounds(ctx context.Context) ([]model.Sound, error) {
	var sounds []model.Sound

	err := r.conn.EvalTyped(`
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
	`, nil, &sounds)
	if err != nil {
		return nil, fmt.Errorf("select random sounds from tarantool storage: %w", err)
	}

	return sounds, nil
}
