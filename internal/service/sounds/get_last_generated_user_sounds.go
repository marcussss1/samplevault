package sounds

import (
	"context"
	"fmt"
	"sort"

	"github.com/marcussss1/simplevault/internal/model"
)

func (s Service) GetLastGeneratedUserSounds(ctx context.Context, userID string) ([]model.Sound, error) {
	sounds, err := s.tarantoolRepository.GetLastGeneratedUserSounds(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get last generated sounds from tarantool repository: %w", err)
	}

	return filterSounds(sounds), nil
}

func filterSounds(sounds []model.Sound) []model.Sound {
	var filteredSounds []model.Sound

	for _, sound := range sounds {
		if sound.IsGenerated == true {
			filteredSounds = append(filteredSounds, sound)
		}
	}

	sort.Slice(filteredSounds, func(i, j int) bool {
		return filteredSounds[i].CreatedAt < filteredSounds[j].CreatedAt
	})

	return filteredSounds[:min(len(filteredSounds), 5)]
}
