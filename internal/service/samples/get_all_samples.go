package samples

import (
	"context"
	"github.com/marcussss1/simplevault/internal/model"
)

func (s Service) GetAllSamples(ctx context.Context, userID string) ([]model.Sample, error) {
	return s.tarantoolRepository.GetAllSamples(ctx, userID)
}
