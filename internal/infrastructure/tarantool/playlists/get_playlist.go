package playlists

import (
	"context"
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
	"github.com/tarantool/go-tarantool"
)

func (r Repository) GetPlaylist(ctx context.Context, id string) (model.Playlist, error) {
	var playlists []model.Playlist

	err := r.conn.SelectTyped("playlists", "primary", 0, 1,
		tarantool.IterEq, tarantool.StringKey{id}, &playlists,
	)
	if err != nil {
		return model.Playlist{}, fmt.Errorf("select playlist from tarantool storage: %w", err)
	}

	if len(playlists) == 0 {
		return model.Playlist{}, nil
	}

	return playlists[0], nil
}
