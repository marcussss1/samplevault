package samples

import (
	"context"
	"github.com/marcussss1/simplevault/internal/model"
)

type tarantoolRepository interface {
	GetAllSamples(ctx context.Context, userID string) ([]model.Sample, error)
}
