package sounds

import (
	"context"
	"fmt"
	"github.com/google/uuid"
)

func (r Repository) LikeSoundByID(ctx context.Context, authorID, soundID string) error {
	// todo транзакция
	_, err := r.conn.Replace("likes", []interface{}{
		uuid.NewString(),
		authorID,
		soundID,
	})
	if err != nil {
		return fmt.Errorf("insert like from tarantool storage: %w", err)
	}

	query := `
		local arg = {...} 
		local id = arg[1]
		
		return box.space.sounds:update(id, {
			{'+', 'likes', 1},
		})
	`

	_, err = r.conn.Eval(query, []interface{}{
		soundID,
	})
	if err != nil {
		return fmt.Errorf("update sound likes from tarantool storage: %w", err)
	}

	return nil
}
