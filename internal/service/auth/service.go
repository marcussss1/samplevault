package auth

import "fmt"

type Service struct {
	tarantoolRepository tarantoolRepository
}

//nolint:goerr113
func NewService(
	tarantoolRepository tarantoolRepository,
) (*Service, error) {
	if tarantoolRepository == nil {
		return nil, fmt.Errorf("tarantoolRepository is nil")
	}

	return &Service{
		tarantoolRepository: tarantoolRepository,
	}, nil
}
