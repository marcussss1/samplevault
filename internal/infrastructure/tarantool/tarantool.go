package tarantool

import (
	"github.com/tarantool/go-tarantool"
)

type Repository struct {
	conn *tarantool.Connection
}

func NewRepository(conn *tarantool.Connection) *Repository {
	return &Repository{
		conn: conn,
	}
}
