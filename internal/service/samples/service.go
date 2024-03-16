package samples

type Service struct {
	tarantoolRepository tarantoolRepository
	minioRepository     minioRepository
}

//nolint:goerr113
func NewService(
	tarantoolRepository tarantoolRepository,
	minioRepository minioRepository,
) *Service {
	return &Service{
		tarantoolRepository: tarantoolRepository,
		minioRepository:     minioRepository,
	}
}
