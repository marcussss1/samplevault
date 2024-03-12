package main

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tarantool/go-tarantool"
	"net/http"
)

type config struct {
	TarantoolDSN      string `env:"TARANTOOL_DSN,required"`
	TarantoolUser     string `env:"TARANTOOL_USER,required"`
	TarantoolPassword string `env:"TARANTOOL_PASSWORD,required"`
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

	e.GET("/api/v1", func(ctx echo.Context) error {
		resp, err := conn.Select("bands", "primary", 0, 1, tarantool.IterEq, []interface{}{3})
		if err != nil {
			fmt.Println(err)
			return err
		}

		fmt.Println(resp)
		fmt.Println("VERSION 3 REQUEST")

		return ctx.JSON(http.StatusOK, struct {
			Message string `json:"message"`
		}{
			Message: "VERSION 2 HANDLE",
		})
	})

	e.GET("/api/v1/health", func(ctx echo.Context) error {
		return ctx.NoContent(http.StatusOK)
	})

	e.GET("/api/v1/read", func(ctx echo.Context) error {
		return ctx.NoContent(http.StatusOK)
	})

	// TODO COPY [/bin/bash] DOCKERFILE

	e.Logger.Fatal(e.Start(":8000"))
}
