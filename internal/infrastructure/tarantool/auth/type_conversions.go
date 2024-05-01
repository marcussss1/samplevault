package auth

import (
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
	"github.com/tarantool/go-tarantool"
)

func toID(resp *tarantool.Response) string {
	var ids []string
	for _, tuples := range resp.Tuples() {
		for _, tuple := range tuples {
			id := tuple.([]interface{})
			ids = append(ids, fmt.Sprint(id[0]))
		}
	}

	if len(ids) == 0 {
		return ""
	}

	return ids[0]
}

func toUser(resp *tarantool.Response) model.User {
	var users []model.User
	for _, tuples := range resp.Tuples() {
		for _, tuple := range tuples {
			user := tuple.([]interface{})
			users = append(users, model.User{
				ID:       fmt.Sprint(user[0]),
				Username: fmt.Sprint(user[1]),
			})
		}
	}

	if len(users) == 0 {
		return model.User{}
	}

	return users[0]
}
