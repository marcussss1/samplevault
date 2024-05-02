package sounds

import "fmt"

type Service struct {
	tarantoolRepository tarantoolRepository
	mlClient            mlClient
	minioRepository     minioRepository
}

//nolint:goerr113
func NewService(
	tarantoolRepository tarantoolRepository,
	mlClient mlClient,
	minioRepository minioRepository,
) (*Service, error) {
	if tarantoolRepository == nil {
		return nil, fmt.Errorf("tarantoolRepository is nil")
	}
	if mlClient == nil {
		return nil, fmt.Errorf("mlClient is nil")
	}
	if minioRepository == nil {
		return nil, fmt.Errorf("minioRepository is nil")
	}

	return &Service{
		tarantoolRepository: tarantoolRepository,
		mlClient:            mlClient,
		minioRepository:     minioRepository,
	}, nil
}
