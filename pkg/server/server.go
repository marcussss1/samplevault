package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	samplescontrollerv1 "github.com/marcussss1/simplevault/internal/controller/http/samplesv1"
	authservice "github.com/marcussss1/simplevault/internal/service/auth"
	httpmiddleware "github.com/marcussss1/simplevault/pkg/middleware"
	"github.com/marcussss1/simplevault/pkg/router"
	"strings"
)

type Server struct {
	E *echo.Echo
}

func NewServer(samplesControllerV1 *samplescontrollerv1.Controller, authService *authservice.Service) *Server {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:80", "http://localhost:8000", "http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST"},
		AllowCredentials: true,
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
	e.Use(httpmiddleware.Auth(authService))
	e = router.Route(e, samplesControllerV1)

	return &Server{
		E: e,
	}
}

func (s Server) Serve() error {
	return s.E.Start(":8000")
}
