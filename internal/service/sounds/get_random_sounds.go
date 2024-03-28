package sounds

import (
	"context"
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
	"math/rand"
)

func (s Service) GetRandomSounds(ctx context.Context) ([]model.Sound, error) {
	sounds, err := s.tarantoolRepository.GetRandomSounds(ctx)
	if err != nil {
		return nil, fmt.Errorf("get all sounds from tarantool repository: %w", err)
	}

	return filterRandomSounds(sounds), nil
}

func filterRandomSounds(sounds []model.Sound) []model.Sound {
	rand.Shuffle(len(sounds), func(i, j int) {
		sounds[i], sounds[j] = sounds[j], sounds[i]
	})

	return sounds[:min(len(sounds), 5)]
}
