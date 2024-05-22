package sounds

import (
	"context"
	"fmt"
)

func (s Service) DislikeSoundByID(ctx context.Context, authorID, soundID string) error {
	err := s.tarantoolRepository.DislikeSoundByID(ctx, authorID, soundID)
	if err != nil {
		return fmt.Errorf("dislike sound by id from tarantool repository: %w", err)
	}

	return nil
}
