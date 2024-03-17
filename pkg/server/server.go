package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	samplescontrollerv1 "github.com/marcussss1/simplevault/internal/controller/http/samplesv1"
	"github.com/marcussss1/simplevault/pkg/router"
	"strings"
)

type Server struct {
	E *echo.Echo
}

func NewServer(samplesControllerV1 *samplescontrollerv1.Controller) *Server {
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
	e = router.Route(e, samplesControllerV1)

	return &Server{
		E: e,
	}
}

func (s Server) Serve() error {
	return s.E.Start(":8000")
}
