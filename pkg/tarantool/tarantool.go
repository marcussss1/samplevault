package tarantool

import (
	"fmt"
	"github.com/tarantool/go-tarantool"
)

func NewClient(cfg Config) (*tarantool.Connection, error) {
	conn, err := tarantool.Connect(cfg.DSN, tarantool.Opts{
		User: cfg.User,
		Pass: cfg.Password,
	})
	if err != nil {
		return nil, fmt.Errorf("tarantool connect: %w", err)
	}

	_, err = conn.Ping()
	if err != nil {
		return nil, fmt.Errorf("tarantool ping: %w", err)
	}

	return conn, nil
}
