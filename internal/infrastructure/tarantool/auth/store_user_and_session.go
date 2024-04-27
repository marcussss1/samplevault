package auth

import (
	"context"
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
)

func (r Repository) StoreUserAndSession(ctx context.Context, user model.FullUser) error {
	// todo транзакция

	_, err := r.conn.Insert("users", []interface{}{
		user.ID,
		user.Username,
		user.Password,
	})
	if err != nil {
		return fmt.Errorf("insert user from tarantool storage: %w", err)
	}

	_, err = r.conn.Insert("sessions", []interface{}{
		user.ID,
		user.SessionID,
	})
	if err != nil {
		return fmt.Errorf("insert session from tarantool storage: %w", err)
	}

	return nil
}
