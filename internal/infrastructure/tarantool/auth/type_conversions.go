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

func toGetUserBySessionID(resp *tarantool.Response) model.GetUserBySessionID {
	var users []model.GetUserBySessionID
	for _, tuples := range resp.Tuples() {
		for _, tuple := range tuples {
			user := tuple.([]interface{})
			users = append(users, model.GetUserBySessionID{
				ID:       fmt.Sprint(user[0]),
				Username: fmt.Sprint(user[1]),
			})
		}
	}

	if len(users) == 0 {
		return model.GetUserBySessionID{}
	}

	return users[0]
}
