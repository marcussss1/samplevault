package sounds

import (
	"context"
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
)

func (s Service) GetLikedSounds(ctx context.Context, userID string) ([]model.Sound, error) {
	sounds, err := s.tarantoolRepository.GetLikedSounds(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get liked sounds from tarantool repository: %w", err)
	}

	return sounds, nil
}
