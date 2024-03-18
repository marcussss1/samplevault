package tarantool

import (
	"context"
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
)

func (r Repository) StoreSample(ctx context.Context, sample model.Sample) error {
	_, err := r.conn.Insert("samples", sample)
	if err != nil {
		return fmt.Errorf("insert sample from tarantool storage: %w", err)
	}

	return nil
}
