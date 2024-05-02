package auth

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/marcussss1/simplevault/internal/model"
)

func (s Service) Signup(ctx context.Context, signupUser model.SignupUserRequest) (model.SignupUserResponse, error) {
	// todo валидация
	// todo проверка на существование юзера по юзернейму
	user := model.FullUser{
		ID:        uuid.NewString(),
		SessionID: uuid.NewString(),
		Username:  signupUser.Username,
		Password:  signupUser.Password,
	}
	err := s.tarantoolRepository.StoreUserAndSession(ctx, user)
	if err != nil {
		return model.SignupUserResponse{}, fmt.Errorf("store user from tarantool repository: %w", err)
	}

	return model.SignupUserResponse{
		ID:        user.ID,
		SessionID: user.SessionID,
		Username:  user.Username,
	}, nil
}
