package auth

import (
	"context"
	"errors"
	"github.com/marcussss1/simplevault/internal/model"
)

func (s Service) Signup(ctx context.Context, signupUser model.SignupUser) (model.User, error) {
	return model.User{}, errors.New("not implemented")
	//	todo валидация
	// todo проверка на существование юзера по юзернейму
	//user := model.FullUser{
	//	ID:        uuid.NewString(),
	//	SessionID: uuid.NewString(),
	//	Username:  signupUser.Username,
	//	Password:  signupUser.Password,
	//}
	//err := s.tarantoolRepository.StoreUserAndSession(ctx, user)
	//if err != nil {
	//	return model.User{}, fmt.Errorf("store user from tarantool repository: %w", err)
	//}
	//
	//return model.User{
	//	ID:        user.ID,
	//	SessionID: user.SessionID,
	//	Username:  user.Username,
	//}, nil
}
