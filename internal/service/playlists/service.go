package playlists

import "fmt"

type Service struct {
	tarantoolRepository  tarantoolRepository
	tarantoolRepository2 tarantoolRepository2
}

//nolint:goerr113
func NewService(
	tarantoolRepository tarantoolRepository,
	tarantoolRepository2 tarantoolRepository2,
) (*Service, error) {
	if tarantoolRepository == nil {
		return nil, fmt.Errorf("tarantoolRepository is nil")
	}
	if tarantoolRepository2 == nil {
		return nil, fmt.Errorf("tarantoolRepository2 is nil")
	}

	return &Service{
		tarantoolRepository:  tarantoolRepository,
		tarantoolRepository2: tarantoolRepository2,
	}, nil
}
