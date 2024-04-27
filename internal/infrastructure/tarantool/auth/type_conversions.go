package auth

import (
	"fmt"
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
