package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET"},
		AllowCredentials: false,
	}))

	e.GET("/api/v1", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, struct {
			Message string `json:"message"`
		}{
			Message: "VERSION 1 HANDLE",
		})
	})

	e.GET("/api/v2", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, struct {
			Message string `json:"message"`
		}{
			Message: "VERSION 2 HANDLE",
		})
	})

	e.Logger.Fatal(e.Start(":8000"))
}
