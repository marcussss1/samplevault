package sounds

import (
	"context"
	"fmt"
)

func (s Service) LikeSoundByID(ctx context.Context, authorID, soundID string) error {
	err := s.tarantoolRepository.LikeSoundByID(ctx, authorID, soundID)
	if err != nil {
		return fmt.Errorf("like sound by id from tarantool repository: %w", err)
	}

	return nil
}
