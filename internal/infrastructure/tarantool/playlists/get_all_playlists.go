package playlists

import (
	"context"
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
	"github.com/tarantool/go-tarantool"
)

func (r Repository) GetAllPlaylists(ctx context.Context) ([]model.Playlist, error) {
	var playlists []model.Playlist

	err := r.conn.SelectTyped("playlists", "primary", 0, 50,
		tarantool.IterAll, tarantool.StringKey{""}, &playlists,
	)
	if err != nil {
		return nil, fmt.Errorf("select all playlists from tarantool storage: %w", err)
	}

	return playlists, nil
}
