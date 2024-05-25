package app

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	samplescontrollerv1 "github.com/marcussss1/simplevault/internal/controller/http/samplesv1"
	miniorepository "github.com/marcussss1/simplevault/internal/infrastructure/minio/sounds"
	"github.com/marcussss1/simplevault/internal/infrastructure/services/ml"
	tarantoolrepository3 "github.com/marcussss1/simplevault/internal/infrastructure/tarantool/auth"
	tarantoolrepository2 "github.com/marcussss1/simplevault/internal/infrastructure/tarantool/playlists"
	tarantoolrepository "github.com/marcussss1/simplevault/internal/infrastructure/tarantool/sounds"
	filesservice "github.com/marcussss1/simplevault/internal/service/audio"
	authservice "github.com/marcussss1/simplevault/internal/service/auth"
	playlistsservice "github.com/marcussss1/simplevault/internal/service/playlists"
	samplesservice "github.com/marcussss1/simplevault/internal/service/sounds"
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

	tarantoolRepository2, err := tarantoolrepository2.NewRepository(tarantoolClient)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	tarantoolRepository3, err := tarantoolrepository3.NewRepository(tarantoolClient)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	mlClient, err := ml.NewClient("https://3c82-178-208-76-164.ngrok-free.app/")
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	samplesService, err := samplesservice.NewService(tarantoolRepository, mlClient, minioRepository)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	filesService, err := filesservice.NewService(minioRepository, tarantoolRepository)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	playlistsService, err := playlistsservice.NewService(tarantoolRepository2, tarantoolRepository)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	authService, err := authservice.NewService(tarantoolRepository3)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	samplesControllerV1, err := samplescontrollerv1.NewController(samplesService, filesService, playlistsService, authService)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	server := server.NewServer(samplesControllerV1, authService)

	err = server.Serve()
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
