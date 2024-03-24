package sounds

import (
	"context"
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
)

func (s Service) GetRandomSounds(ctx context.Context) ([]model.Sound, error) {
	sounds, err := s.tarantoolRepository.GetRandomSounds(ctx)
	if err != nil {
		return nil, fmt.Errorf("get random sounds from tarantool repository: %w", err)
	}

	return sounds, nil
}
