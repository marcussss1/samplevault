package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tarantool/go-tarantool"
	"net/http"
)

func main() {
	fmt.Println()
	fmt.Println("VERSION 3 APP")
	fmt.Println()

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET"},
		AllowCredentials: false,
	}))

	conn, err := tarantool.Connect("tarantool:3301", tarantool.Opts{
		User: "tarantool",
		Pass: "tarantool",
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
