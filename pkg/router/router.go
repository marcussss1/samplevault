package router

import (
	"github.com/labstack/echo/v4"
	samplescontrollerv1 "github.com/marcussss1/simplevault/internal/controller/http/samplesv1"
	"net/http"
)

func Route(e *echo.Echo, ctrl *samplescontrollerv1.Controller) *echo.Echo {
	e.GET("/api/v1/sounds", ctrl.GetAllSounds)
	e.GET("/api/v1/sounds/generate", ctrl.GenerateSound)
	e.GET("/api/v1/sounds/download/:filename", ctrl.DownloadSound)
	e.POST("/api/v1/sounds/upload", ctrl.UploadSound)
	e.GET("/api/v1/sounds/random", ctrl.GetRandomSounds)
	e.GET("/api/v1/sounds/last_generated", ctrl.GetLastGeneratedUserSounds)

	e.GET("/api/v1/playlists/random", ctrl.GetRandomPlaylists)
	e.GET("/api/v1/playlists/:playlist_id", ctrl.GetPlaylist)

	e.GET("/api/v1/health", func(ctx echo.Context) error {
		return ctx.NoContent(http.StatusOK)
	})
	e.GET("/api/v1/read", func(ctx echo.Context) error {
		return ctx.NoContent(http.StatusOK)
	})

	return e
}
