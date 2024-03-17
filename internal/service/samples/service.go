package samples

import "fmt"

type Service struct {
	tarantoolRepository tarantoolRepository
	minioRepository     minioRepository
}

//nolint:goerr113
func NewService(
	tarantoolRepository tarantoolRepository,
	minioRepository minioRepository,
) (*Service, error) {
	if tarantoolRepository == nil {
		return nil, fmt.Errorf("tarantoolRepository is nil")
	}
	if minioRepository == nil {
		return nil, fmt.Errorf("minioRepository is nil")
	}

	return &Service{
		tarantoolRepository: tarantoolRepository,
		minioRepository:     minioRepository,
	}, nil
}
