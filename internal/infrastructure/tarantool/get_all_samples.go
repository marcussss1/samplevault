package tarantool

import (
	"context"
	"github.com/marcussss1/simplevault/internal/model"
	"github.com/tarantool/go-tarantool"
)

//nolint:gomnd
func (r Repository) GetAllSamples(ctx context.Context, userID string) ([]model.Sample, error) {
	var samples []model.Sample

	err := r.conn.SelectTyped(
		"samples",
		"primary",
		0,
		10,
		tarantool.IterEq,
		tarantool.StringKey{userID},
		&samples,
	)
	if err != nil {
		return nil, err
	}

	return samples, nil
}
