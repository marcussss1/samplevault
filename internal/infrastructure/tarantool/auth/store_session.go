package auth

import (
	"context"
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
)

func (r Repository) StoreSession(ctx context.Context, session model.Session) error {
	_, err := r.conn.Insert("sessions", []interface{}{
		session.ID,
		session.UserID,
	})
	if err != nil {
		return fmt.Errorf("insert session from tarantool storage: %w", err)
	}

	return nil
}
