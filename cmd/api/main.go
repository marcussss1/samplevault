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

	e.GET("/", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, struct {
			Message string `json:"message"`
		}{
			Message: "/ HANDLE",
		})
	})

	e.GET("/api", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, struct {
			Message string `json:"message"`
		}{
			Message: "API HANDLE",
		})
	})

	e.Logger.Fatal(e.Start(":8000"))
}
