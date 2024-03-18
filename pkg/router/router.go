package router

import (
	"github.com/labstack/echo/v4"
	samplescontrollerv1 "github.com/marcussss1/simplevault/internal/controller/http/samplesv1"
	"net/http"
)

func Route(e *echo.Echo, ctrl *samplescontrollerv1.Controller) *echo.Echo {
	e.GET("/api/v1/samples", ctrl.GetAllSamples)
	e.GET("/api/v1/samples/generate", ctrl.GenerateSample)
	e.GET("/api/v1/samples/download/:filename", ctrl.DownloadSample)
	e.POST("/api/v1/samples/upload", ctrl.UploadSample)

	e.GET("/api/v1/health", func(ctx echo.Context) error {
		return ctx.NoContent(http.StatusOK)
	})
	e.GET("/api/v1/read", func(ctx echo.Context) error {
		return ctx.NoContent(http.StatusOK)
	})

	return e
}
