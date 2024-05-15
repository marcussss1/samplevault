package sounds

import (
	"context"
	"fmt"
)

func (s Service) DeleteSoundByID(ctx context.Context, id string) error {
	err := s.tarantoolRepository.DeleteSoundByID(ctx, id)
	if err != nil {
		return fmt.Errorf("delete sound by id from tarantool repository: %w", err)
	}

	return nil
}
