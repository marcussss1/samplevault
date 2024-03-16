package main

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tarantool/go-tarantool"
	"io"
	"net/http"
	"os"
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
		AllowMethods:     []string{"GET"},
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

	e.POST("/api/v1/samples/generate", func(ctx echo.Context) error {
		file, err := os.Open("sample.mp3")
		if err != nil {
			return err
		}
		defer file.Close()

		ctx.Response().Header().Set(echo.HeaderContentType, "audio/mpeg")

		// Отправляем файл клиенту
		_, err = io.Copy(ctx.Response().Writer, file)
		if err != nil {
			return err
		}

		return ctx.NoContent(http.StatusOK)

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

		//return ctx.JSON(http.StatusOK, []Sample{
		//	{
		//		ID:                uuid.MustParse("a2802d62-b006-4949-8fa0-07328bd26719"),
		//		AuthorID:          uuid.MustParse("a2802d62-b006-4949-8fa0-07328bd26719"),
		//		AudioURL:          "Ссылка на аудио",
		//		IconURL:           "Ссылка на иконку",
		//		Title:             "Название сэмпла",
		//		Duration:          "Длительность",
		//		MusicalInstrument: "Kick",
		//		Genre:             "Hip Hop",
		//		IsFavourite:       false,
		//	},
		//})
	})
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
