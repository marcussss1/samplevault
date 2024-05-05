package auth

import (
	"context"
	"fmt"
)

func (s Service) DeleteSessionByID(ctx context.Context, sessionID string) error {
	err := s.tarantoolRepository.DeleteSessionByID(ctx, sessionID)
	if err != nil {
		return fmt.Errorf("delete session by id from tarantool repository: %w", err)
	}

	return nil
}
