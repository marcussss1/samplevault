package auth

import (
	"context"
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
)

func (r Repository) DeleteSessionByID(ctx context.Context, sessionID string) error {
	query := `
		local arg = {...} 
		local sessionID = arg[1]
		return box.space.sessions.index.id:delete(sessionID)
	`

	resp, err := r.conn.Eval(query, []interface{}{
		sessionID,
	})
	if err != nil {
		return fmt.Errorf("delete session by id from tarantool storage: %w", err)
	}

	fmt.Println("len: ", len(resp.Tuples()))

	if len(resp.Tuples()) == 0 {
		return model.ErrNotFound
	}

	return nil
}
