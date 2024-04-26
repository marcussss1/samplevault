package auth

import (
	"context"
	"github.com/marcussss1/simplevault/internal/model"
)

type tarantoolRepository interface {
	StoreUser(ctx context.Context, user model.FullUser) error
}
