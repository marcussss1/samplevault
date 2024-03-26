package playlists

import (
	"context"
	"github.com/marcussss1/simplevault/internal/model"
)

type tarantoolRepository interface {
	GetAllPlaylists(ctx context.Context) ([]model.Playlist, error)
	GetPlaylist(ctx context.Context, id string) (model.Playlist, error)
}

type tarantoolRepository2 interface {
	GetSound(ctx context.Context, id string) (model.Sound, error)
}
