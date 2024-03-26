package sounds

import (
	"context"
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
	"github.com/tarantool/go-tarantool"
)

//nolint:gomnd
func (r Repository) GetSound(ctx context.Context, id string) (model.Sound, error) {
	var sounds []model.Sound

	err := r.conn.SelectTyped("sounds", "primary", 0, 1,
		tarantool.IterEq, tarantool.StringKey{id}, &sounds,
	)
	if err != nil {
		return model.Sound{}, fmt.Errorf("select sound from tarantool storage: %w", err)
	}

	if len(sounds) == 0 {
		return model.Sound{}, nil
	}

	return sounds[0], nil
}
