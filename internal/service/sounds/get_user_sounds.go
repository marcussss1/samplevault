package sounds

import (
	"context"
	"fmt"

	"github.com/marcussss1/simplevault/internal/model"
)

func (s Service) GetUserSounds(ctx context.Context, userID string) ([]model.Sound, error) {
	sounds, err := s.tarantoolRepository.GetUserSounds(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get user sounds from tarantool repository: %w", err)
	}

	return sounds, nil
}
