package playlists

import (
	"context"
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
	"math/rand"
)

func (s Service) GetRandomPlaylists(ctx context.Context) ([]model.RandomPlaylist, error) {
	playlists, err := s.tarantoolRepository.GetAllPlaylists(ctx)
	if err != nil {
		return nil, fmt.Errorf("get all playlists from tarantool repository: %w", err)
	}

	return filterRandomPlaylists(playlists), nil
}

func filterRandomPlaylists(playlists []model.Playlist) []model.RandomPlaylist {
	var filteredPlaylists []model.RandomPlaylist

	for _, playlist := range playlists {
		filteredPlaylists = append(filteredPlaylists, model.RandomPlaylist{
			ID:      playlist.ID,
			IconURL: playlist.IconURL,
		})
	}

	rand.Shuffle(len(filteredPlaylists), func(i, j int) {
		filteredPlaylists[i], filteredPlaylists[j] = filteredPlaylists[j], filteredPlaylists[i]
	})

	return filteredPlaylists[:min(len(filteredPlaylists), 5)]
}
