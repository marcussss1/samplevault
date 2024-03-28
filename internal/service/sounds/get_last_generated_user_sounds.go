package sounds

import (
	"context"
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
)

func (s Service) GetLastGeneratedUserSounds(ctx context.Context, userID string) ([]model.Sound, error) {
	sounds, err := s.tarantoolRepository.GetLastGeneratedUserSounds(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get last generated user sounds from tarantool repository: %w", err)
	}

	return sounds, nil
}

//return box.space.sounds.index.primary:select({}, {limit = 5, iterator = 'ALL', order = 'desc'})
