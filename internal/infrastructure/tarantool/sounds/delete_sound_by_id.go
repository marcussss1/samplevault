package sounds

import (
	"context"
	"fmt"
)

//nolint:gomnd
func (r Repository) DeleteSoundByID(ctx context.Context, id string) error {
	_, err := r.conn.Delete("sounds", "primary", []interface{}{id})
	if err != nil {
		return fmt.Errorf("delete sound by id: %w", err)
	}

	return nil
}
