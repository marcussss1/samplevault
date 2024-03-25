package sounds

import (
	"context"
	"fmt"

	"github.com/marcussss1/simplevault/internal/model"
)

func (s Service) GetAllSounds(ctx context.Context) ([]model.Sound, error) {
	sounds, err := s.tarantoolRepository.GetAllSounds(ctx)
	if err != nil {
		return nil, fmt.Errorf("get all sounds from tarantool repository: %w", err)
	}

	return sounds, nil
}
