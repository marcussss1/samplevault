package audio

import "fmt"

type Service struct {
	minioRepository     minioRepository
	tarantoolRepository tarantoolRepository
}

//nolint:goerr113
func NewService(
	minioRepository minioRepository,
	tarantoolRepository tarantoolRepository,
) (*Service, error) {
	if minioRepository == nil {
		return nil, fmt.Errorf("minioRepository is nil")
	}
	if tarantoolRepository == nil {
		return nil, fmt.Errorf("tarantoolRepository is nil")
	}

	return &Service{
		minioRepository:     minioRepository,
		tarantoolRepository: tarantoolRepository,
	}, nil
}
