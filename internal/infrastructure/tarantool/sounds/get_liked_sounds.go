package sounds

import (
	"context"
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
)

func (r Repository) GetLikedSounds(ctx context.Context, userID string) ([]model.Sound, error) {
	query := `
		local resp = {}
		local arg = {...} 
		local authorID = arg[1]

		local likes = box.space.likes.index.author_id:select(authorID)
		local sounds = box.space.sounds:select() 
		for _, like in ipairs(likes) do
			for _, sound in ipairs(sounds) do
				if like[3] == sound[1] then
					table.insert(resp, sound)
				end
			end
		end
	
		return resp`

	resp, err := r.conn.Eval(query, []interface{}{
		userID,
	})
	if err != nil {
		return nil, fmt.Errorf("select liked sounds from tarantool storage: %w", err)
	}

	return toSounds(resp), nil
}
