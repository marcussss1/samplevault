package main

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/marcussss1/simplevault/internal/controller/http/samplesv1"
	minio2 "github.com/marcussss1/simplevault/internal/infrastructure/minio"
	tarantool2 "github.com/marcussss1/simplevault/internal/infrastructure/tarantool"
	"github.com/marcussss1/simplevault/internal/service/samples"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/tarantool/go-tarantool"
	"net/http"
	"strings"
)

// TODO без дефолта
type config struct {
	TarantoolDSN      string `env:"TARANTOOL_DSN,required" envDefault:"tarantool:3301"`
	TarantoolUser     string `env:"TARANTOOL_USER,required" envDefault:"tarantool"`
	TarantoolPassword string `env:"TARANTOOL_PASSWORD,required" envDefault:"tarantool"`

	MinioDSN       string `env:"MINIO_DSN,required" envDefault:"minio:9000"`
	MinioAccessKey string `env:"MINIO_ACCESS_KEY,required" envDefault:"your_access_key"`
	MinioSecretKey string `env:"MINIO_SECRET_KEY,required" envDefault:"your_secret_key"`
}

func main() {
	var cfg config
	if err := env.Parse(&cfg); err != nil {
		fmt.Println("parse ", err)
		return
	}

	conn, err := tarantool.Connect(cfg.TarantoolDSN, tarantool.Opts{
		User: cfg.TarantoolUser,
		Pass: cfg.TarantoolPassword,
	})
	if err != nil {
		fmt.Println("connect ", err)
		return
	}
	defer conn.Close()

	_, err = conn.Ping()
	if err != nil {
		fmt.Println("ping  ", err)
		return
	}

	minioClient, err := minio.New(cfg.MinioDSN, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.MinioAccessKey, cfg.MinioSecretKey, ""),
		Secure: false,
	})
	if err != nil {
		panic(err)
	}

	minioRepository := minio2.NewRepository(minioClient)
	tarantoolRepository := tarantool2.NewRepository(conn)
	samplesService := samples.NewService(tarantoolRepository, minioRepository)
	samplesController := samplesv1.NewController(samplesService)

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowCredentials: false,
	}))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: func(ctx echo.Context) bool {
			url := ctx.Request().URL.String()
			if strings.Contains(url, "health") || strings.Contains(url, "read") {
				return true
			}
			return false
		},
	}))

	e.GET("/api/v1/samples", samplesController.GetAllSamples)
	e.GET("/api/v1/samples/generate", samplesController.GenerateSample)
	e.GET("/api/v1/samples/download", samplesController.DownloadSample)
	e.POST("/api/v1/samples/upload", samplesController.UploadSample)

	e.GET("/api/v1/health", func(ctx echo.Context) error {
		return ctx.NoContent(http.StatusOK)
	})

	e.GET("/api/v1/read", func(ctx echo.Context) error {
		return ctx.NoContent(http.StatusOK)
	})

	// TODO COPY [/bin/bash] DOCKERFILE

	e.Logger.Fatal(e.Start(":8000"))
}
