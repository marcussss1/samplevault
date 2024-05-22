package sounds

import (
	"context"
	"fmt"
)

func (r Repository) DislikeSoundByID(ctx context.Context, authorID, soundID string) error {
	// todo транзакция
	query := `
		local arg = {...} 
		local authorID = arg[1]
		local soundID = arg[2]
		local records = box.space.likes.index.sound_id:select(soundID)
		
		for _, tuple in ipairs(records) do
			if tuple[2] == authorID then
				box.space.likes.index.primary:delete(tuple[1])
				break
			end
		end


		return box.space.sounds:update(soundID, {
			{'-', 'likes', 1},
		})
	`

	_, err := r.conn.Eval(query, []interface{}{
		authorID,
		soundID,
	})
	if err != nil {
		return fmt.Errorf("update sound likes from tarantool storage: %w", err)
	}

	return nil
}
