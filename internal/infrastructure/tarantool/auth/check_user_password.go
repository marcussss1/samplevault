package auth

import (
	"context"
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
)

func (r Repository) CheckUserPassword(ctx context.Context, loginUser model.LoginUser) (model.User, error) {
	query := `
		local arg = {...} 
		local username = arg[1]
		local password = arg[2]

		user = box.space.users.index.username:select(username)
        if user then
			if user.password ~= password then
				return {}
			else
         		return {
					user.ID
				}
			end
    	else
        	return {}
    	end	
	`

	resp, err := r.conn.Eval(query, []interface{}{
		loginUser.Username,
		loginUser.Password,
	})
	if err != nil {
		return model.User{}, fmt.Errorf("select user from tarantool storage: %w", err)
	}

	id := toID(resp)
	if id == "" {
		return model.User{}, model.ErrNotFound
	}

	return model.User{
		ID:       id,
		Username: loginUser.Username,
	}, nil
}
