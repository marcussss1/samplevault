package app

import (
	"github.com/caarlos0/env/v6"
	samplescontrollerv1 "github.com/marcussss1/simplevault/internal/controller/http/samplesv1"
	miniorepository "github.com/marcussss1/simplevault/internal/infrastructure/minio"
	tarantoolrepository "github.com/marcussss1/simplevault/internal/infrastructure/tarantool"
	samplesservice "github.com/marcussss1/simplevault/internal/service/samples"
	"github.com/marcussss1/simplevault/pkg/minio"
	"github.com/marcussss1/simplevault/pkg/server"
	"github.com/marcussss1/simplevault/pkg/tarantool"
	"log"
)

type Config struct {
	Tarantool tarantool.Config
	Minio     minio.Config
}

func Run() error {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return err
	}

	tarantoolClient, err := tarantool.NewClient(cfg.Tarantool)
	if err != nil {
		return err
	}
	defer tarantoolClient.Close()

	minioClient, err := minio.NewClient(cfg.Minio)
	if err != nil {
		return err
	}

	minioRepository, err := miniorepository.NewRepository(minioClient)
	if err != nil {
		return err
	}

	tarantoolRepository, err := tarantoolrepository.NewRepository(tarantoolClient)
	if err != nil {
		return err
	}

	samplesService, err := samplesservice.NewService(tarantoolRepository, minioRepository)
	if err != nil {
		log.Fatal(err)
	}

	samplesControllerV1, err := samplescontrollerv1.NewController(samplesService)
	if err != nil {
		return err
	}

	server := server.NewServer(samplesControllerV1)

	err = server.Serve()
	if err != nil {
		return err
	}

	return nil
}
