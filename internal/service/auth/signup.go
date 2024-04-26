package auth

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/marcussss1/simplevault/internal/model"
)

func (s Service) Signup(ctx context.Context, signupUser model.SignupUser) (model.User, error) {
	// todo валидация
	user := model.FullUser{
		ID:        uuid.NewString(),
		SessionID: uuid.NewString(),
		Username:  signupUser.Username,
		Password:  signupUser.Password,
	}
	err := s.tarantoolRepository.StoreUser(ctx, user)
	if err != nil {
		return model.User{}, fmt.Errorf("upload sound from minio repository: %w", err)
	}

	return model.User{
		ID:       user.ID,
		Username: user.Username,
	}, nil
}
