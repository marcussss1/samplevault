package middleware

import (
	"context"
	"github.com/marcussss1/simplevault/internal/model"
)

type authService interface {
	GetUserBySessionID(ctx context.Context, sessionID string) (model.GetUserBySessionID, error)
}
