package auth

import (
	"context"
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
)

func (r Repository) CheckExistUsername(ctx context.Context, username string) error {
	query := `
		local arg = {...} 
		local username = arg[1]
		return box.space.users.index.username:select(username)
	`

	resp, err := r.conn.Eval(query, []interface{}{
		username,
	})
	if err != nil {
		return fmt.Errorf("select user from tarantool storage: %w", err)
	}

	if len(resp.Tuples()) != 0 {
		return model.ErrAlreadyExist
	}

	return nil
}
