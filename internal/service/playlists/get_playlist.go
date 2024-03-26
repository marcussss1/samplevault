package playlists

import (
	"context"
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
)

func (s Service) GetPlaylist(ctx context.Context, id string) (model.FullPlaylist, error) {
	playlist, err := s.tarantoolRepository.GetPlaylist(ctx, id)
	if err != nil {
		return model.FullPlaylist{}, fmt.Errorf("get playlist from tarantool repository: %w", err)
	}

	var sounds []model.Sound
	for _, soundID := range playlist.SoundIDs {
		sound, err := s.tarantoolRepository2.GetSound(ctx, soundID)
		if err != nil {
			return model.FullPlaylist{}, fmt.Errorf("get sound from tarantool repository: %w", err)
		}

		sounds = append(sounds, sound)
	}

	return model.FullPlaylist{
		ID:          playlist.ID,
		AuthorID:    playlist.AuthorID,
		IconURL:     playlist.IconURL,
		CreatedAt:   playlist.CreatedAt,
		Title:       playlist.Title,
		Description: playlist.Description,
		Sounds:      sounds,
	}, nil
}
