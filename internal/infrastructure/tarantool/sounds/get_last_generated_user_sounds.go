package sounds

import (
	"context"
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
	"github.com/tarantool/go-tarantool"
)

//nolint:gomnd
func (r Repository) GetLastGeneratedUserSounds(ctx context.Context, userID string) ([]model.Sound, error) {
	var sounds []model.Sound

	err := r.conn.SelectTyped("sounds", "author_id", 0, 20,
		tarantool.IterEq, tarantool.StringKey{userID}, &sounds,
	)
	if err != nil {
		return nil, fmt.Errorf("select last generated user sounds from tarantool storage: %w", err)
	}

	return sounds, nil
}
