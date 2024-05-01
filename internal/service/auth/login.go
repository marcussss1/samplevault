package auth

import (
	"context"
	"errors"
	"github.com/marcussss1/simplevault/internal/model"
)

func (s Service) Login(ctx context.Context, loginUser model.LoginUser) (model.User, error) {
	return model.User{}, errors.New("not implemented")
	//// todo валидация
	//user, err := s.tarantoolRepository.CheckUserPassword(ctx, loginUser)
	//if err != nil {
	//	return model.User{}, fmt.Errorf("check user password from tarantool repository: %w", err)
	//}
	//
	//user.SessionID = uuid.NewString()
	//err = s.tarantoolRepository.StoreSession(ctx, model.Session{
	//	ID:     user.SessionID,
	//	UserID: user.ID,
	//})
	//if err != nil {
	//	return model.User{}, fmt.Errorf("store session from tarantool repository: %w", err)
	//}
	//
	//return user, nil
}
