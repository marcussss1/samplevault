package sounds

import (
	"context"
	"github.com/marcussss1/simplevault/internal/model"
)

type tarantoolRepository interface {
	GetAllSounds(ctx context.Context) ([]model.Sound, error)
}
