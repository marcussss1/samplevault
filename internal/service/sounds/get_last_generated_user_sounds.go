package sounds

import (
	"context"
	"fmt"
	"sort"

	"github.com/marcussss1/simplevault/internal/model"
)

func (s Service) GetLastGeneratedUserSounds(ctx context.Context, userID string) ([]model.Sound, error) {
	sounds, err := s.tarantoolRepository.GetAllSounds(ctx)
	if err != nil {
		return nil, fmt.Errorf("get all sounds from tarantool repository: %w", err)
	}

	return filterLastGeneratedUserSounds(sounds, userID), nil
}

func filterLastGeneratedUserSounds(sounds []model.Sound, userID string) []model.Sound {
	var filteredSounds []model.Sound

	for _, sound := range sounds {
		fmt.Println(sound.IsGenerated, "  :  ", sound.AuthorID)
		if sound.IsGenerated == true && sound.AuthorID == userID {
			filteredSounds = append(filteredSounds, sound)
		}
	}

	sort.Slice(filteredSounds, func(i, j int) bool {
		return filteredSounds[i].CreatedAt < filteredSounds[j].CreatedAt
	})

	return filteredSounds[:min(len(filteredSounds), 5)]
}
