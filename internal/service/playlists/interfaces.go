package playlists

import (
	"context"
	"github.com/marcussss1/simplevault/internal/model"
)

type tarantoolRepository interface {
	GetAllPlaylists(ctx context.Context, userID string) ([]model.Playlist, error)
}
