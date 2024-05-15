package sounds

import (
	"context"
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
	"github.com/tarantool/go-tarantool"
)

//nolint:gomnd
func (r Repository) GetAllSounds(ctx context.Context) ([]model.Sound, error) {
	var sounds []model.Sound

	err := r.conn.SelectTyped("sounds", "primary", 0, 1000,
		tarantool.IterAll, tarantool.StringKey{""}, &sounds,
	)
	if err != nil {
		return nil, fmt.Errorf("select all sounds from tarantool storage: %w", err)
	}

	return sounds, nil
}
