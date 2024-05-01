package auth

import (
	"context"
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
)

func (s Service) GetUserBySessionID(ctx context.Context, sessionID string) (model.User, error) {
	user, err := s.tarantoolRepository.GetUserBySessionID(ctx, sessionID)
	if err != nil {
		return model.User{}, fmt.Errorf("get user by session id from tarantool repository: %w", err)
	}

	return user, nil
}
