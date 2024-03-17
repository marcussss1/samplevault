package samples

import (
	"context"
	"fmt"

	"github.com/marcussss1/simplevault/internal/model"
)

func (s Service) GetAllSamples(ctx context.Context, userID string) ([]model.Sample, error) {
	samples, err := s.tarantoolRepository.GetAllSamples(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get all samples from tarantool repository: %w", err)
	}

	return samples, nil
}
