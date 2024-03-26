package sounds

import (
	"context"
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
	"github.com/tarantool/go-tarantool"
)

//nolint:gomnd
func (r Repository) GetSound(ctx context.Context, id string) (model.Sound, error) {
	var sound model.Sound

	err := r.conn.SelectTyped("sounds", "primary", 0, 1,
		tarantool.IterEq, tarantool.StringKey{id}, &sound,
	)
	if err != nil {
		return model.Sound{}, fmt.Errorf("select sound from tarantool storage: %w", err)
	}

	return sound, nil
}
