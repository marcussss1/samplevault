package playlists

import (
	"context"
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
)

func (s Service) GetAllPlaylists(ctx context.Context, userID string) ([]model.Playlist, error) {
	playlists, err := s.tarantoolRepository.GetAllPlaylists(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get all playlists from tarantool repository: %w", err)
	}

	return playlists, nil
}
