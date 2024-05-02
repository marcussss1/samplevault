package auth

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/marcussss1/simplevault/internal/model"
)

func (s Service) Login(ctx context.Context, loginUser model.LoginUserRequest) (model.LoginUserResponse, error) {
	// todo валидация
	user, err := s.tarantoolRepository.CheckUserPassword(ctx, loginUser)
	if err != nil {
		return model.LoginUserResponse{}, fmt.Errorf("check user password from tarantool repository: %w", err)
	}

	user.SessionID = uuid.NewString()
	err = s.tarantoolRepository.StoreSession(ctx, model.Session{
		ID:     user.SessionID,
		UserID: user.ID,
	})
	if err != nil {
		return model.LoginUserResponse{}, fmt.Errorf("store session from tarantool repository: %w", err)
	}

	return user, nil
}
