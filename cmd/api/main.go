package main

import (
	"context"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/tarantool/go-tarantool"
	"io"
	"net/http"
)

// TODO без дефолта
type config struct {
	TarantoolDSN      string `env:"TARANTOOL_DSN,required" envDefault:"tarantool:3301"`
	TarantoolUser     string `env:"TARANTOOL_USER,required" envDefault:"tarantool"`
	TarantoolPassword string `env:"TARANTOOL_PASSWORD,required" envDefault:"tarantool"`
}

type Sample struct {
	ID                string `json:"id"`
	AuthorID          string `json:"author_id"`
	AudioURL          string `json:"audio_url"`
	IconURL           string `json:"icon_url"`
	Title             string `json:"title"`
	Duration          string `json:"duration"`
	MusicalInstrument string `json:"musical_instrument"`
	Genre             string `json:"genre"`
	IsFavourite       bool   `json:"is_favourite"`
}

func main() {
	fmt.Println()
	fmt.Println("VERSION 5 APP")
	fmt.Println()

	var cfg config
	if err := env.Parse(&cfg); err != nil {
		fmt.Println("parse ", err)
		return
	}

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowCredentials: false,
	}))
	e.Use(middleware.Logger())

	conn, err := tarantool.Connect(cfg.TarantoolDSN, tarantool.Opts{
		User: cfg.TarantoolUser,
		Pass: cfg.TarantoolPassword,
	})
	if err != nil {
		fmt.Println("connect ", err)
		return
	}
	defer conn.Close()

	fmt.Println("CONNECT SUCCESS")

	_, err = conn.Ping()
	if err != nil {
		fmt.Println("ping  ", err)
		return
	}

	fmt.Println("PING SUCCESS")

	e.GET("/api/v1/samples", func(ctx echo.Context) error {
		var samples []Sample

		err := conn.SelectTyped(
			"samples",
			"primary",
			0, 2,
			tarantool.IterEq,
			tarantool.StringKey{"a2802d62-b006-4949-8fa0-07328bd26719"},
			&samples,
		)
		if err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, samples)
	})

	endpoint := "minio:9000"
	accessKey := "your_access_key"
	secretKey := "your_secret_key"

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: false,
	})
	if err != nil {
		panic(err)
	}

	// Create bucket
	//err = minioClient.MakeBucket(context.Background(), "samples", minio.MakeBucketOptions{})
	//if err != nil {
	//	log.Fatal(err)
	//}

	e.POST("/api/v1/samples/generate", func(ctx echo.Context) error {
		//var samples []Sample
		//
		//err := conn.SelectTyped(
		//	"samples",
		//	"primary",
		//	0, 2,
		//	tarantool.IterEq,
		//	tarantool.StringKey{"a2802d62-b006-4949-8fa0-07328bd26719"},
		//	&samples,
		//)
		//if err != nil {
		//	return err
		//}
		//
		//return ctx.JSON(http.StatusOK, samples)
		objectReader, err := minioClient.GetObject(context.Background(), "samples", "sample.mp3", minio.GetObjectOptions{})
		if err != nil {
			return err
		}
		defer objectReader.Close()

		// Отправляем содержимое файла по HTTP клиенту
		ctx.Response().Header().Set(echo.HeaderContentType, "application/octet-stream")
		ctx.Response().Header().Set(echo.HeaderContentDisposition, fmt.Sprintf("attachment; filename=%s", "sample.mp3"))

		if _, err := io.Copy(ctx.Response().Writer, objectReader); err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, "ok")
	})

	e.POST("/api/v1/samples/download", func(ctx echo.Context) error {
		objectReader, err := minioClient.GetObject(context.Background(), "samples", "sample.mp3", minio.GetObjectOptions{})
		if err != nil {
			return err
		}
		defer objectReader.Close()

		// Отправляем содержимое файла по HTTP клиенту
		ctx.Response().Header().Set(echo.HeaderContentType, "application/octet-stream")
		ctx.Response().Header().Set(echo.HeaderContentDisposition, fmt.Sprintf("attachment; filename=%s", "sample.mp3"))

		if _, err := io.Copy(ctx.Response().Writer, objectReader); err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, "ok")
		//return ctx.(http.StatusOK)
		//file, err := ctx.FormFile("file")
		//if err != nil {
		//	return err
		//}
		//
		//src, err := file.Open()
		//if err != nil {
		//	return err
		//}
		//defer src.Close()
		//
		//_, err = minioClient.PutObject(context.Background(), "samples", file.Filename, src, file.Size, minio.PutObjectOptions{})
		//if err != nil {
		//	return err
		//}
		//
		//return ctx.NoContent(http.StatusOK)
	})

	e.POST("/api/v1/samples/upload", func(ctx echo.Context) error {
		//file, err := ctx.FormFile("file")
		//if err != nil {
		//	return err
		//}
		//
		//src, err := file.Open()
		//if err != nil {
		//	return err
		//}
		//defer src.Close()
		maxSize := int64(64 << 20)
		err := ctx.Request().ParseMultipartForm(maxSize)
		if err != nil {
			return err
		}

		file, header, err := ctx.Request().FormFile("audio")
		if err != nil {
			return err
		}

		_, err = minioClient.PutObject(context.Background(), "samples", "sample.mp3", file, header.Size, minio.PutObjectOptions{})
		if err != nil {
			return err
		}

		return ctx.NoContent(http.StatusOK)
	})

	//e.POST("/api/v1/samples/generate", func(ctx echo.Context) error {
	//	//file, err := os.Open("/tmp/sample.mp3")
	//	//if err != nil {
	//	//	return err
	//	//}
	//	//defer file.Close()
	//	//
	//	//ctx.Response().Header().Set(echo.HeaderContentType, "audio/mpeg")
	//	//
	//	//// Отправляем файл клиенту
	//	//_, err = io.Copy(ctx.Response().Writer, file)
	//	//if err != nil {
	//	//	return err
	//	//}
	//	//
	//	//return ctx.NoContent(http.StatusOK)
	//
	//	//var samples []Sample
	//	//
	//	//err := conn.SelectTyped(
	//	//	"samples",
	//	//	"primary",
	//	//	0, 2,
	//	//	tarantool.IterEq,
	//	//	tarantool.StringKey{"a2802d62-b006-4949-8fa0-07328bd26719"},
	//	//	&samples,
	//	//)
	//	//if err != nil {
	//	//	return err
	//	//}
	//
	//	//return ctx.JSON(http.StatusOK, []Sample{
	//	//	{
	//	//		ID:                uuid.MustParse("a2802d62-b006-4949-8fa0-07328bd26719"),
	//	//		AuthorID:          uuid.MustParse("a2802d62-b006-4949-8fa0-07328bd26719"),
	//	//		AudioURL:          "Ссылка на аудио",
	//	//		IconURL:           "Ссылка на иконку",
	//	//		Title:             "Название сэмпла",
	//	//		Duration:          "Длительность",
	//	//		MusicalInstrument: "Kick",
	//	//		Genre:             "Hip Hop",
	//	//		IsFavourite:       false,
	//	//	},
	//	//})
	//})
	//return samples, nil

	//fmt.Println(resp)
	//fmt.Println("VERSION 3 REQUEST")

	//	return ctx.JSON(http.StatusOK, struct {
	//		Message string `json:"message"`
	//	}{
	//		Message: "VERSION 2 HANDLE",
	//	})
	//})

	e.GET("/api/v1/health", func(ctx echo.Context) error {
		return ctx.NoContent(http.StatusOK)
	})

	e.GET("/api/v1/read", func(ctx echo.Context) error {
		return ctx.NoContent(http.StatusOK)
	})

	// TODO COPY [/bin/bash] DOCKERFILE

	e.Logger.Fatal(e.Start(":8000"))
}
