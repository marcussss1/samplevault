package sounds

import (
	"context"
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
)

func (s Service) EditSoundByID(ctx context.Context, sound model.EditSound) error {
	err := s.tarantoolRepository.EditSoundByID(ctx, sound)
	if err != nil {
		return fmt.Errorf("edit sound by id from tarantool repository: %w", err)
	}

	return nil
}
