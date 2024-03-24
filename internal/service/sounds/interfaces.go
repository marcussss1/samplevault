package sounds

import (
	"context"
	"github.com/marcussss1/simplevault/internal/model"
)

type tarantoolRepository interface {
	GetUserSounds(ctx context.Context, userID string) ([]model.Sound, error)
	GetRandomSounds(ctx context.Context) ([]model.Sound, error)
}
