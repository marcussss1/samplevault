package sounds

import (
	"fmt"

	"github.com/tarantool/go-tarantool"
)

type Repository struct {
	conn *tarantool.Connection
}

//nolint:goerr113
func NewRepository(conn *tarantool.Connection) (*Repository, error) {
	if conn == nil {
		return nil, fmt.Errorf("conn is nil")
	}

	return &Repository{
		conn: conn,
	}, nil
}
