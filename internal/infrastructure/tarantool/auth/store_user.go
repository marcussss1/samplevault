package auth

import (
	"context"
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
)

func (r Repository) StoreUser(ctx context.Context, user model.FullUser) error {
	// todo транзакция

	_, err := r.conn.Insert("users", []interface{}{
		user.ID,
		user.Username,
		user.Password,
	})
	if err != nil {
		return fmt.Errorf("insert sound from tarantool storage: %w", err)
	}

	_, err = r.conn.Insert("sessions", []interface{}{
		user.ID,
		user.SessionID,
	})

	return nil
}
