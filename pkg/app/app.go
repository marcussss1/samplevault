package app

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	samplescontrollerv1 "github.com/marcussss1/simplevault/internal/controller/http/samplesv1"
	miniorepository "github.com/marcussss1/simplevault/internal/infrastructure/minio"
	tarantoolrepository "github.com/marcussss1/simplevault/internal/infrastructure/tarantool"
	filesservice "github.com/marcussss1/simplevault/internal/service/files"
	samplesservice "github.com/marcussss1/simplevault/internal/service/samples"
	"github.com/marcussss1/simplevault/pkg/minio"
	"github.com/marcussss1/simplevault/pkg/server"
	"github.com/marcussss1/simplevault/pkg/tarantool"
)

type Config struct {
	Tarantool tarantool.Config
	Minio     minio.Config
}

func Run() error {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return fmt.Errorf("%w", err)
	}

	tarantoolClient, err := tarantool.NewClient(cfg.Tarantool)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	defer tarantoolClient.Close()

	minioClient, err := minio.NewClient(cfg.Minio)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	minioRepository, err := miniorepository.NewRepository(minioClient)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	tarantoolRepository, err := tarantoolrepository.NewRepository(tarantoolClient)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	samplesService, err := samplesservice.NewService(tarantoolRepository)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	filesService, err := filesservice.NewService(minioRepository, tarantoolRepository)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	samplesControllerV1, err := samplescontrollerv1.NewController(samplesService, filesService)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	server := server.NewServer(samplesControllerV1)

	err = server.Serve()
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
