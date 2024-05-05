package auth

import (
	"context"
	"github.com/marcussss1/simplevault/internal/model"
)

type tarantoolRepository interface {
	StoreUserAndSession(ctx context.Context, user model.FullUser) error
	StoreSession(ctx context.Context, session model.Session) error
	CheckUserPassword(ctx context.Context, loginUser model.LoginUserRequest) (model.LoginUserResponse, error)
	CheckExistUsername(ctx context.Context, username string) error
	GetUserBySessionID(ctx context.Context, sessionID string) (model.GetUserBySessionID, error)
	DeleteSessionByID(ctx context.Context, sessionID string) error
}
