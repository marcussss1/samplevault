package auth

import (
	"context"
	"github.com/marcussss1/simplevault/internal/model"
)

type tarantoolRepository interface {
	StoreUserAndSession(ctx context.Context, user model.FullUser) error
	StoreSession(ctx context.Context, session model.Session) error
	CheckUserPassword(ctx context.Context, loginUser model.LoginUser) (model.User, error)
	GetUserBySessionID(ctx context.Context, sessionID string) (model.User, error)
}
