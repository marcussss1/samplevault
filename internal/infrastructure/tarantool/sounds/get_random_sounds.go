package sounds

import (
	"context"
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
	"github.com/tarantool/go-tarantool"
)

func (r Repository) GetRandomSounds(ctx context.Context) ([]model.Sound, error) {
	var sounds []model.Sound

	err := r.conn.SelectTyped("sounds", "primary", 0, 5,
		tarantool.IterAll, nil, &sounds,
	)
	if err != nil {
		return nil, fmt.Errorf("select sounds from tarantool storage: %w", err)
	}

	return sounds, nil
}
