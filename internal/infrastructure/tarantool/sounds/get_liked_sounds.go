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
		
		likes = box.space.likes.index.author_id:select(authorID)
		for _, like in ipairs(likes) do
	    	local sound = box.space.sounds.index.primary:select(like[3])
			table.insert(resp, sound)
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
